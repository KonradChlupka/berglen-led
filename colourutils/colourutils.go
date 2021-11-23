package colourutils

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
