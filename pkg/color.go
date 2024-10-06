package goray

import (
	"fmt"
	"math"
)

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

func (c Color) ToPpm() string {
	r := int(math.Ceil(clamp(255.999 * c.x)))
	g := int(math.Ceil(clamp(255.999 * c.y)))
	b := int(math.Ceil(clamp(255.999 * c.z)))

	return fmt.Sprintf("%d %d %d", r, g, b)
}

func clamp(x float64) float64 {
	if x < 0.0 {
		return 0.0
	} else if x > 255.0 {
		return 255.0
	}
	return x
}
