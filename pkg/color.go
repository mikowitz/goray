package goray

type Color = Tuple

func NewColor(r, g, b float64) Color {
	return Color{x: r, y: g, z: b, w: 0}
}

func (c Color) Prod(d Color) Color {
	return NewColor(
		c.x*d.x,
		c.y*d.y,
		c.z*d.z,
	)
}
