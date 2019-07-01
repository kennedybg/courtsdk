package courtsdk

import (
	"strings"
	"sync"
	"time"
)

//ConfigMutex used to lock config map before read/write
var ConfigMutex sync.Mutex

func getDefaultElasticConfig() map[string]interface{} {
	ConfigMutex.Lock()
	config := map[string]interface{}{
		"URL":                  GetEnvString("ELASTIC_URL", "http://localhost"),
		"Port":                 GetEnvInt("ELASTIC_PORT", 9200),
		"Index":                GetEnvString("ELASTIC_INDEX", "jurisprudences_dev"),
		"RetryConnectionDelay": GetEnvInt("ELASTIC_RETRY_CONNECTION_DELAY", 10),
		"RetryPingDelay":       GetEnvInt("ELASTIC_RETRY_PING_DELAY", 5),
	}
	ConfigMutex.Unlock()
	return config
}

func getDefaultEngineConfig() map[string]interface{} {
	ConfigMutex.Lock()
	config := map[string]interface{}{
		"IsAsync":             strings.ToUpper(GetEnvString("ENGINE_IS_ASYNC", "TRUE")) == "TRUE",
		"MaxFailures":         GetEnvInt("ENGINE_MAX_FAILURES", 25),
		"RequestsPerInterval": GetEnvInt("ENGINE_REQUESTS_PER_INTERVAL", 10),
		"RequestDelay":        time.Duration(GetEnvInt("ENGINE_REQUEST_DELAY", 3500)),
		"RequestTimeout":      time.Duration(GetEnvInt("ENGINE_REQUEST_TIMEOUT", 25)),
		"GoRoutineRange":      GetEnvInt("ENGINE_GOROUTINE_RANGE", 200),
		"MaxRecoveries":       GetEnvInt("ENGINE_MAX_RECOVERIES", 5),
	}
	ConfigMutex.Unlock()
	return config
}

func getDefaultControlConfig() map[string]interface{} {
	ConfigMutex.Lock()
	config := map[string]interface{}{
		"IsConcurrent":         strings.ToUpper(GetEnvString("CONTROL_IS_CONCURRENT", "FALSE")) == "TRUE",
		"MaxConcurrentEngines": GetEnvInt("CONTROL_MAX_CONCURRENT_ENGINES", 2),
		"LastGoRoutineRange":   -1,
		"ActionDelay":          time.Duration(GetEnvInt("CONTROL_ACTION_DELAY", 25)),
	}
	ConfigMutex.Unlock()
	return config
}

//ElasticConfig - config
var ElasticConfig = getDefaultElasticConfig()

//EngineConfig - Config engine behavior
var EngineConfig = getDefaultEngineConfig()

//ControlConfig - Config the control behavior
var ControlConfig = getDefaultControlConfig()

//ResetConfigsToDefault use to reset all configs to the default values.
func ResetConfigsToDefault() {
	ElasticConfig = getDefaultElasticConfig()
	EngineConfig = getDefaultEngineConfig()
	ControlConfig = getDefaultControlConfig()
}
