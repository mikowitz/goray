package goray

import (
	"math"
)

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

func (ray Ray) Intersect(sphere Sphere) Intersections {
	ray2 := ray.Transform(sphere.Transform.Inverse())
	sphereToRay := ray2.Origin.Sub(NewPoint(0, 0, 0))

	a := ray2.Direction.Dot(ray2.Direction)
	b := 2 * ray2.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - (4.0 * a * c)

	if discriminant < 0.0 {
		return Intersections{}
	}

	return Intersections{
		NewIntersection((-b-math.Sqrt(discriminant))/(2.0*a), sphere),
		NewIntersection((-b+math.Sqrt(discriminant))/(2.0*a), sphere),
	}
}

func (ray Ray) Transform(m Matrix) Ray {
	return NewRay(
		m.Mult(ray.Origin),
		m.Mult(ray.Direction),
	)
}
