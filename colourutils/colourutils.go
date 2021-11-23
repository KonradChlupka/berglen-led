package colourutils

import (
	"strconv"
	"strings"

	"github.com/hsluv/hsluv-go"
)

const (
	OFF   uint32 = 0x000000
	WHITE uint32 = 0xffffff
	RED   uint32 = 0xff0000
	GREEN uint32 = 0x00ff00
	BLUE  uint32 = 0x0000ff
)

func RGBToColour(r uint8, g uint8, b uint8) uint32 {
	return uint32(uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

func CreateHSVRainbowSlice() ([]uint32, error) {
	const len = 360
	s := make([]uint32, len)

	for i := 0; i < len; i++ {
		h := hsluv.HsluvToHex(float64(i), 100, 80)
		n, err := strconv.ParseUint(strings.TrimPrefix(h, "#"), 16, 32)
		if err != nil {
			return nil, err
		}

		s[i] = uint32(n)
	}

	return s, nil
}
