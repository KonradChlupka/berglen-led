package engine

import "github.com/KonradChlupka/berglen-led/colourutils"

type Christmas struct {
	leds     LEDEngine
	ledCount int

	ticker int
}

func (c *Christmas) RenderFrame() error {
	singleColour := func(colour uint32) {
		for i := 0; i < c.ledCount; i++ {
			c.leds.Leds(0)[i] = colour
		}
	}

	if c.ticker < 10 {
		// Colour all GREEN.
		singleColour(colourutils.GREEN)
	} else if c.ticker < 20 {
		// Colour all GREEN.
		singleColour(colourutils.RED)
	}

	c.ticker++
	if c.ticker >= 20 {
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

	return c, nil
}
