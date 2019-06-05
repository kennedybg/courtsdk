package courtsdk

import "log"

// NewControl creates a new Control instance with default configuration.
func NewControl(options ...func(*Control)) *Control {
	return &Control{}
}

//Register - register a new engine
func (control *Control) Register(engine *Engine) {
	isValid, errorMessage := validateEngine(engine)
	if isValid {
		control.Engines = append(control.Engines, *engine)
		log.Println("[ENGINE] COURT ->", engine.Court, " BASE ->", engine.Base, " successfully registered.")
		return
	}
	log.Println(errorMessage)
}

func validateEngine(engine *Engine) (bool, string) {
	if engine.Court == "" {
		return false, "[ENGINE] Failed an Engine must have a COURT defined."
	}
	if engine.Base == "" {
		return false, "[ENGINE] Failed an Engine must have a BASE defined."
	}
	if engine.EntryPoint == nil {
		return false, "[ENGINE] Failed an Engine must have a ENTRYPOINT defined."
	}
	return true, ""
}
