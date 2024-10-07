package goray

import "math"

type Tuple struct {
	x, y, z, w float64
}
type Point = Tuple
type Vector = Tuple

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x: x, y: y, z: z, w: w}
}

func NewPoint(x, y, z float64) Point {
	return NewTuple(x, y, z, 1)
}

func NewVector(x, y, z float64) Point {
	return NewTuple(x, y, z, 0)
}

func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

func (u Tuple) Add(v Tuple) Tuple {
	return NewTuple(u.x+v.x, u.y+v.y, u.z+v.z, u.w+v.w)
}

func (u Tuple) Sub(v Tuple) Tuple {
	return NewTuple(u.x-v.x, u.y-v.y, u.z-v.z, u.w-v.w)
}

func (u Tuple) Neg() Tuple {
	return NewTuple(-u.x, -u.y, -u.z, -u.w)
}

func (u Tuple) Mul(t float64) Tuple {
	return NewTuple(u.x*t, u.y*t, u.z*t, u.w*t)
}

func (u Tuple) Div(t float64) Tuple {
	return u.Mul(1 / t)
}

func (u Tuple) Magnitude() float64 {
	return math.Sqrt(u.x*u.x + u.y*u.y + u.z*u.z + u.w*u.w)
}

func (u Tuple) Normalize() Tuple {
	return u.Div(u.Magnitude())
}

func (u Tuple) Dot(v Tuple) float64 {
	return u.x*v.x + u.y*v.y + u.z*v.z + u.w*v.w
}

func (u Tuple) Cross(v Tuple) Tuple {
	return NewVector(
		u.y*v.z-u.z*v.y,
		u.z*v.x-u.x*v.z,
		u.x*v.y-u.y*v.x,
	)
}

func (u Tuple) Reflect(n Tuple) Tuple {
	return u.Sub(n.Mul(2.0 * u.Dot(n)))
}
