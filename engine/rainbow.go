package engine

import (
	"context"

	"github.com/KonradChlupka/berglen-led/colourutils"
)

// RainbowRGB starts an RGB rainbow pattern, and only stops once the context is cancelled.
func (l *lightstrip) RainbowRGB(ctx context.Context) error {

	s, err := colourutils.CreateHSVRainbowSlice()
	if err != nil {
		return err
	}

	ledCount := l.Length()

	for {
		for i := 0; i < ledCount; i++ {
			l.leds.Leds(0)[i] = s[i]
		}

		if err := l.leds.Render(); err != nil {
			return err
		}

		s = shiftUint32Slice(s)

		if err := l.leds.Wait(); err != nil {
			return err
		}
	}

	return nil
}

func shiftUint32Slice(s []uint32) []uint32 {
	x := s[0] // get the 0 index element from slice
	s = s[1:]
	s = append(s, x)
	return s
}
