package courtsdk

import (
	"strings"
	"time"
)

//EngineConfig - Config engine behavior
var EngineConfig = map[string]interface{}{
	"isAsync":             strings.ToUpper(GetEnvString("ENGINE_IS_ASYNC", "TRUE")) == "TRUE",
	"MaxFailures":         GetEnvInt("ENGINE_MAX_FAILURES", 25),
	"RequestsPerInterval": GetEnvInt("ENGINE_REQUESTS_PER_INTERVAL", 10),
	"RequestDelay":        time.Duration(GetEnvInt("ENGINE_REQUEST_DELAY", 3500)),
	"RequestTimeout":      time.Duration(GetEnvInt("ENGINE_REQUEST_TIMEOUT", 25)),
	"GoRoutineRange":      GetEnvInt("ENGINE_GOROUTINE_RANGE", 200),
	"MaxRecoveries":       GetEnvInt("ENGINE_MAX_RECOVERIES", 5),
}
