package courtsdk

import (
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/olivere/elastic"
)

// NewEngine creates a new Engine instance with default configuration
func NewEngine(options ...func(*Engine)) *Engine {
	engine := &Engine{}
	engine.Collector = GetDefaultcollector()
	engine.ResponseChannel = make(chan int)
	engine.PageSize = 1
	var wg sync.WaitGroup
	engine.Lock = &wg
	for _, attr := range options {
		attr(engine)
	}
	return engine
}

// Court set the current court.
func Court(court string) func(*Engine) {
	return func(engine *Engine) {
		engine.Court = court
	}
}

// Base set the current base.
func Base(base string) func(*Engine) {
	return func(engine *Engine) {
		engine.Base = base
	}
}

// Start set the start index
func Start(start int) func(*Engine) {
	return func(engine *Engine) {
		if engine.IsConcurrent {
			ControlConfig["LastGoRoutineRange"] = start - 1
		}
		engine.Start = start
	}
}

// End set the end index
func End(end int) func(*Engine) {
	return func(engine *Engine) {
		engine.End = end
	}
}

// PageSize set the pagination size.
func PageSize(pageSize int) func(*Engine) {
	return func(engine *Engine) {
		engine.PageSize = pageSize
	}
}

// Collector set the engine private collector (colly)
func Collector(collector *colly.Collector) func(*Engine) {
	return func(engine *Engine) {
		engine.Collector = collector
	}
}

// ElasticClient set the engine private elasticsearch client
func ElasticClient(client *elastic.Client) func(*Engine) {
	return func(engine *Engine) {
		engine.ElasticClient = client
	}
}

// EntryPoint set a function to start the engine.
func EntryPoint(entry func(engine *Engine)) func(*Engine) {
	return func(engine *Engine) {
		engine.EntryPoint = entry
	}
}

// ResponseChannel set an own private channel for the engine.
func ResponseChannel(responseChannel chan int) func(*Engine) {
	return func(engine *Engine) {
		engine.ResponseChannel = responseChannel
	}
}

// Lock set a private WaitGroup for the engine.
func Lock(lock *sync.WaitGroup) func(*Engine) {
	return func(engine *Engine) {
		engine.Lock = lock
	}
}

// Concurrency set how many replicas and range (both greater than zero).
func Concurrency(maxReplicas int, replicaRange int) func(*Engine) {
	return func(engine *Engine) {
		if maxReplicas > 0 && replicaRange > 0 {
			engine.IsConcurrent = true
			engine.MaxReplicas = maxReplicas
			engine.ReplicaRange = replicaRange
		}
	}
}

//InitElastic - Initialize an Elasticsearch client with Elastic configs.
func (engine *Engine) InitElastic() {
	var err error
	elasticFullURL := ElasticConfig["URL"].(string) + ":" + strconv.Itoa(ElasticConfig["Port"].(int))
	engine.ElasticClient, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(elasticFullURL))
	if err != nil {
		log.Println("[FAILED] Connect to Elasticsearch.", err)
		log.Println("[WARNING] Retrying in ", strconv.Itoa(ElasticConfig["RetryConnectionDelay"].(int)), " seconds...")
		time.Sleep(time.Duration(ElasticConfig["RetryConnectionDelay"].(int)) * time.Second)
		engine.InitElastic()
		return
	}
	engine.pingElasticSearch(elasticFullURL)
}

func (engine *Engine) pingElasticSearch(elasticFullURL string) {
	context, cancelContext := GetNewContext()
	defer cancelContext()
	info, code, err := engine.ElasticClient.Ping(elasticFullURL).Do(context)
	if err != nil {
		log.Println("[FAILED] Ping to Elasticsearch.", err)
		log.Println("[WARNING] Retrying in ", strconv.Itoa(ElasticConfig["RetryPingDelay"].(int)), " seconds...")
		time.Sleep(time.Duration(ElasticConfig["RetryPingDelay"].(int)) * time.Second)
		engine.InitElastic()
		return
	}
	log.Printf("[SUCCESS] Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

//ConnectedToIndex - Check if the given index exist.
func (engine *Engine) ConnectedToIndex() bool {
	index := ElasticConfig["Index"].(string)
	context, cancelContext := GetNewContext()
	defer cancelContext()
	exists, err := engine.ElasticClient.IndexExists(index).Do(context)
	if err != nil {
		log.Println("[FAILED] Unable to connect to index -> ["+index+"]", err)
		return false
	}
	if !exists {
		log.Println("[WARNING] Index -> [" + index + "] not found. Attempting to create...")
		createIndex, err := engine.ElasticClient.CreateIndex(index).BodyString(GetElasticMapping()).Do(context)
		if err != nil {
			log.Println("[FAILED] Create index -> ["+index+"].", err)
			return false
		}
		if !createIndex.Acknowledged {
			log.Println("[WARNING] Index -> [" + index + "] was created, but not acknowledged.")
			return false
		}
		log.Println("[SUCCESS] Index -> [" + index + "] was created and acknowledged.")
		return true
	}
	log.Println("[SUCCESS] Index -> [" + index + "] was found, sending data to it...")
	return true
}

//Persist - send data to Elasticsearch.
func (engine *Engine) Persist(jurisprudence Jurisprudence) {
	uid := jurisprudence.Court + "-" + engine.Base + "-" + jurisprudence.DocumentID
	context, cancelContext := GetNewContext()
	defer cancelContext()
	_, err := engine.ElasticClient.Index().
		Index(ElasticConfig["Index"].(string)).
		Type("_doc").
		Id(uid).
		BodyJson(jurisprudence).
		Do(context)
	if err != nil {
		log.Println("[FAILED][CREATE] Save document ["+jurisprudence.DocumentID+"]["+jurisprudence.DocumentType+"]:", err)
		engine.ResponseChannel <- http.StatusInternalServerError
	}
	engine.ResponseChannel <- http.StatusOK
}

//GetDocumentType - returns the document type.
func (engine *Engine) GetDocumentType() string {
	switch engine.Base {
	case "baseAcordaos":
		return "Acórdãos"
	case "baseSumulas":
		return "Súmulas"
	case "baseSumulasVinculantes":
		return "Súmulas Vinculantes"
	case "basePresidencia":
		return "Decisões da Presidência"
	case "baseRepercussao":
		return "Repercussão Geral"
	default:
		return "CUSTOM[" + engine.Base + "]"
	}
}

func (engine *Engine) runAsSequential() {
	engine.InitElastic()
	if engine.ConnectedToIndex() {
		engine.Recoveries = 0
		for engine.Recoveries < EngineConfig["MaxRecoveries"].(int) {
			engine.EntryPoint(engine)
			if engine.Done {
				engine.logSuccess()
				return
			}
			engine.logFailure()
			engine.setRecoveryStart()
			time.Sleep(ControlConfig["ActionDelay"].(time.Duration) * time.Second)
			engine.Failures = 0
			engine.Recoveries++
		}
	}
}

func (engine *Engine) runAsConcurrent() {
	activeEngines := 0
	maxEngines := engine.MaxReplicas
	activeEnginesChannel := make(chan int)
	maxEnginesChannel := make(chan int)
	elasticMutex := sync.Mutex{}
	for {
		if activeEngines == 0 && maxEngines == 0 {
			return
		}
		select {
		case value := <-activeEnginesChannel:
			activeEngines += value
		case value := <-maxEnginesChannel:
			maxEngines += value
		default:
			if activeEngines < maxEngines {
				activeEngines++
				go engine.spawnEngine(activeEnginesChannel, maxEnginesChannel, &elasticMutex)
			}
		}
	}
}

func (engine Engine) spawnEngine(activeEnginesChannel chan int, maxEnginesChannel chan int, elasticMutex *sync.Mutex) {
	engine.InitElastic()
	elasticMutex.Lock()
	connectedToIndex := engine.ConnectedToIndex()
	elasticMutex.Unlock()
	if connectedToIndex {
		engine.setRange()
		for engine.Recoveries < EngineConfig["MaxRecoveries"].(int) {
			engine.EntryPoint(&engine)
			if engine.Done {
				engine.logSuccess()
				activeEnginesChannel <- -1
				return
			}
			engine.logFailure()
			engine.setRecoveryStart()
			time.Sleep(ControlConfig["ActionDelay"].(time.Duration) * time.Second)
			engine.Failures = 0
			engine.Recoveries++
		}
		maxEnginesChannel <- -1
	}
	activeEnginesChannel <- -1
}

//SetRange - set a valid range for an engine
func (engine *Engine) setRange() {
	lastRange := ControlConfig["LastGoRoutineRange"].(int)
	engine.Start = lastRange + 1
	engine.End = lastRange + engine.ReplicaRange
	if lastRange < engine.End {
		ControlConfig["LastGoRoutineRange"] = engine.End
	}
	log.Println("[INFO] New Engine replica RANGE: ", engine.Start, " to", engine.End)
}

func (engine *Engine) setRecoveryStart() {
	if engine.CurrentIndex != 0 {
		engine.Start = engine.CurrentIndex - (engine.Failures * engine.PageSize)
	}
}

func (engine Engine) logSuccess() {
	str := "[ENGINE] COURT -> " + engine.Court
	str += " BASE -> " + engine.Base + " ended successfully."
	log.Println(str)
}

func (engine Engine) logFailure() {
	str := "[ENGINE] COURT -> " + engine.Court
	str += " BASE -> " + engine.Base + " " + strconv.Itoa(EngineConfig["MaxFailures"].(int)) + " times."
	str += " Last ID requested: " + strconv.Itoa(engine.CurrentIndex)
	str += " Trying to recover from index -> " + strconv.Itoa(engine.Start)
	log.Println(str)
}
