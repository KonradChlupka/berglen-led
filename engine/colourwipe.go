package engine

import "github.com/KonradChlupka/berglen-led/colourutils"

type ColourWipe struct {
	leds LEDEngine

	started      bool
	currentIndex int
	colour       uint32

	ledCount int
}

func (cw *ColourWipe) RenderFrame() error {
	if !cw.started {
		for i := 0; i < cw.ledCount; i++ {
			cw.leds.Leds(0)[i] = colourutils.OFF
		}
		cw.started = true
	} else {
		cw.leds.Leds(0)[cw.currentIndex] = cw.colour
		cw.currentIndex++
	}

	if err := cw.leds.Render(); err != nil {
		return err
	}

	if err := cw.leds.Wait(); err != nil {
		return err
	}

	return nil
}

func (cw *ColourWipe) IsDone() bool {
	return cw.currentIndex >= cw.ledCount
}

// ColourWipe wipes the strip of LEDs with a constant colour.
func (l *lightstrip) ColourWipe(colour uint32) (TemporaryLEDProgram, error) {
	colourWipe := &ColourWipe{
		leds:     l.leds,
		colour:   colour,
		ledCount: len(l.leds.Leds(0)),
	}

	return colourWipe, nil
}
