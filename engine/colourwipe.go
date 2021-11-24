package engine

type ColourWipe struct {
	leds LEDEngine

	currentIndex int
	colour       uint32

	ledCount int
}

func (cw *ColourWipe) RenderFrame() error {
	cw.leds.Leds(0)[cw.currentIndex] = cw.colour
	cw.currentIndex++

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
