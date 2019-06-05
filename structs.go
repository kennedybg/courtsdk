package courtsdk

import (
	"sync"

	"github.com/gocolly/colly"
	"github.com/olivere/elastic"
)

//Engine is a structure used for define information of a engine.
type Engine struct {
	Court           string
	Base            string
	Failures        int
	Start           int
	End             int
	PageSize        int
	CurrentIndex    int
	Recoveries      int
	Done            bool
	EntryPoint      func(engine *Engine)
	ResponseChannel chan int
	Collector       *colly.Collector
	ElasticClient   *elastic.Client
	Lock            *sync.WaitGroup
}
