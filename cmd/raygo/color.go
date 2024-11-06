package main

type Color struct {
	R, G, B float64
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R * 0xffff)
	g = uint32(c.G * 0xffff)
	b = uint32(c.B * 0xffff)
	a = 0xffff
	return r, g, b, a
}
