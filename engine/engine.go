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

	RainbowRGB() (LEDProgram, error)
	ColourWipe(uint32) (TemporaryLEDProgram, error)
	Christmas() (LEDProgram, error)
}

// LEDProgram is a simple program that can be used to display things.
type LEDProgram interface {
	RenderFrame() error
}

type TemporaryLEDProgram interface {
	LEDProgram
	IsDone() bool
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

// Length returns the number of LED's in the LED strip.
func (l *lightstrip) Length() int {
	return len(l.leds.Leds(0))
}

// WaitChan returns a channel where one can wait for the current frame to finish rendering.
func (l *lightstrip) Waitchan() <-chan struct{} {
	finished := make(chan struct{})

	go func() {
		l.leds.Wait()
		finished <- struct{}{}
	}()

	return finished
}
