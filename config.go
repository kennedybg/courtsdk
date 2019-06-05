package courtsdk

import (
	"strings"
	"time"
)

//ElasticConfig - config
var ElasticConfig = map[string]interface{}{
	"URL":                  GetEnvString("ELASTIC_URL", "http://localhost"),
	"Port":                 GetEnvInt("ELASTIC_PORT", 9200),
	"Index":                GetEnvString("ELASTIC_INDEX", "jurisprudences_dev"),
	"RetryConnectionDelay": GetEnvInt("ELASTIC_RETRY_CONNECTION_DELAY", 10),
	"RetryPingDelay":       GetEnvInt("ELASTIC_RETRY_PING_DELAY", 5),
}

//EngineConfig - Config engine behavior
var EngineConfig = map[string]interface{}{
	"IsAsync":             strings.ToUpper(GetEnvString("ENGINE_IS_ASYNC", "TRUE")) == "TRUE",
	"MaxFailures":         GetEnvInt("ENGINE_MAX_FAILURES", 25),
	"RequestsPerInterval": GetEnvInt("ENGINE_REQUESTS_PER_INTERVAL", 10),
	"RequestDelay":        time.Duration(GetEnvInt("ENGINE_REQUEST_DELAY", 3500)),
	"RequestTimeout":      time.Duration(GetEnvInt("ENGINE_REQUEST_TIMEOUT", 25)),
	"GoRoutineRange":      GetEnvInt("ENGINE_GOROUTINE_RANGE", 200),
	"MaxRecoveries":       GetEnvInt("ENGINE_MAX_RECOVERIES", 5),
}

//ControlConfig - Config the control behavior
var ControlConfig = map[string]interface{}{
	"IsConcurrent":         strings.ToUpper(GetEnvString("CONTROL_IS_CONCURRENT", "FALSE")) == "TRUE",
	"MaxConcurrentEngines": GetEnvInt("CONTROL_MAX_CONCURRENT_ENGINES", 2),
	"LastGoRoutineRange":   -1,
	"ActionDelay":          time.Duration(GetEnvInt("CONTROL_ACTION_DELAY", 25)),
}
