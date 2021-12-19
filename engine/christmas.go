package engine

import (
	"github.com/KonradChlupka/berglen-led/colourutils"
)

type Christmas struct {
	leds     LEDEngine
	ledCount int

	ticker    int
	candycane []uint32
}

func (c *Christmas) Setup() error {
	// Sections of white and red.
	var flipped bool
	c.candycane = make([]uint32, c.ledCount)
	for i := 0; i < c.ledCount; i++ {
		if i%20 == 0 {
			flipped = !flipped
		}
		if flipped {
			c.candycane[i] = colourutils.WHITE
		} else {
			c.candycane[i] = colourutils.RED
		}
	}

	return nil
}

func (c *Christmas) RenderFrame() error {
	singleColour := func(colour uint32) {
		for i := 0; i < c.ledCount; i++ {
			c.leds.Leds(0)[i] = colour
		}
	}

	if c.ticker < 100 {
		// Colour all GREEN.
		singleColour(colourutils.GREEN)
	} else if c.ticker < 200 {
		// Colour all GREEN.
		singleColour(colourutils.RED)
	} else if c.ticker < 500 {
		// Run a candy cane effect?
		for i := 0; i < c.ledCount; i++ {
			c.leds.Leds(0)[i] = c.candycane[i]
		}
		c.candycane = shiftUint32Slice(c.candycane)
	}

	c.ticker++
	if c.ticker >= 500 {
		c.ticker = 0
	}

	if err := c.leds.Render(); err != nil {
		return err
	}

	return c.leds.Wait()
}

// Christmas starts a new christmas LED program.
func (l *lightstrip) Christmas() (LEDProgram, error) {
	c := &Christmas{
		leds:     l.leds,
		ledCount: l.Length(),
	}

	err := c.Setup()

	return c, err
}
