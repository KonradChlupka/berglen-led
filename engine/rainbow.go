package engine

import (
	"github.com/KonradChlupka/berglen-led/colourutils"
)

type RainbowRGB struct {
	leds LEDEngine

	rainbowSlice []uint32
	ledCount     int
}

func (r *RainbowRGB) Setup() error {
	s, err := colourutils.CreateHSVRainbowSlice()
	if err != nil {
		return err
	}

	r.rainbowSlice = s
	r.ledCount = len(r.leds.Leds(0))

	return nil
}

func (r *RainbowRGB) RenderFrame() error {
	for i := 0; i < r.ledCount; i++ {
		r.leds.Leds(0)[i] = r.rainbowSlice[i]
	}

	if err := r.leds.Render(); err != nil {
		return err
	}

	rendered := make(chan error)
	go func() {
		rendered <- r.leds.Wait()
	}()

	r.rainbowSlice = shiftUint32Slice(r.rainbowSlice)

	return <-rendered
}

// RainbowRGB starts an RGB rainbow pattern, and only stops once the context is cancelled.
func (l *lightstrip) RainbowRGB() (LEDProgram, error) {
	rgb := &RainbowRGB{
		leds: l.leds,
	}

	err := rgb.Setup()
	if err != nil {
		return nil, err
	}

	return rgb, nil
}

func shiftUint32Slice(s []uint32) []uint32 {
	x := s[0] // get the 0 index element from slice
	s = s[1:]
	s = append(s, x)
	return s
}
