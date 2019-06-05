package courtsdk

// NewEngine creates a new Engine instance with default configuration
func NewEngine(options ...func(*Engine)) *Engine {
	return &Engine{}
}
