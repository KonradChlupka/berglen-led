package colourutils

const (
	WHITE uint32 = 0x000000
	RED   uint32 = 0xff0000
)

func RGBToColour(r uint8, g uint8, b uint8) uint32 {
	return uint32(uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}
