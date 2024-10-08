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

func (ray Ray) Intersect(shape Shape) Intersections {
	ray2 := ray.Transform(shape.GetTransform().Inverse())
	return shape.LocalIntersect(ray2)
}

func (ray Ray) Transform(m Matrix) Ray {
	return NewRay(
		m.Mult(ray.Origin),
		m.Mult(ray.Direction),
	)
}
