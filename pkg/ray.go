package goray

type Ray struct {
	Origin    Point
	Direction Vector
}

func NewRay(origin Point, direction Vector) Ray {
	return Ray{Origin: origin, Direction: direction}
}

func (r Ray) At(t float64) Point {
	return r.Origin.Add(r.Direction.Mul(t))
}
