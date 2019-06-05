package courtsdk

// NewControl creates a new Control instance with default configuration.
func NewControl(options ...func(*Control)) *Control {
	return &Control{}
}
