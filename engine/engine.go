package engine

// LEDEngine talks to the actual LED circuits to display output.
type LEDEngine interface {
	Init() error
	Render() error
	Wait() error
	Fini()
	Leds(channel int) []uint32
}

// Engine is a simple wrapper around the LEDEngine with our custom programs on top.
type Engine interface {
	Init() error
	Close()
	// TODO: Add some actual functionality
}

// lightstrip is our concrete implementation of an LED lightstrip.
type lightstrip struct {
	leds LEDEngine
}

// NewLightstrip instantiates a new lightstrip object.
func NewLightstrip(ledEngine LEDEngine) *lightstrip {
	return &lightstrip{
		leds: ledEngine,
	}
}

// Init implements the Init function for instantiating a connection to the LED strip.
func (l *lightstrip) Init() error {
	return l.leds.Init()
}

// Close gracefully closes the connection to the LED strip.
func (l *lightstrip) Close() {
	l.leds.Fini()
}
