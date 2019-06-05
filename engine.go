package courtsdk

import (
	"sync"

	"github.com/gocolly/colly"
	"github.com/olivere/elastic"
)

// NewEngine creates a new Engine instance with default configuration
func NewEngine(options ...func(*Engine)) *Engine {
	return &Engine{}
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
	return func(eng *Engine) {
		ControlConfig["LastGoRoutineRange"] = start - 1
		eng.Start = start
	}
}

// End set the end index
func End(end int) func(*Engine) {
	return func(eng *Engine) {
		eng.End = end
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
