package engine

// ColourWipe wipes the strip of LEDs with a constant colour.
func (l *lightstrip) ColourWipe(colour uint32) error {
	for i := 0; i < len(l.leds.Leds(0)); i++ {
		l.leds.Leds(0)[i] = colour
		if err := l.leds.Render(); err != nil {
			return err
		}
		if err := l.leds.Wait(); err != nil {
			return err
		}
	}
	return nil
}
