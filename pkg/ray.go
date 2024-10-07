package goray

import "math"

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
	sphereToRay := ray.Origin.Sub(sphere.Center)

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - sphere.Radius*sphere.Radius

	discriminant := b*b - 4.0*a*c

	if discriminant < 0.0 {
		return Intersections{}
	}

	return Intersections{
		NewIntersection((-b-math.Sqrt(discriminant))/2.0*a, sphere),
		NewIntersection((-b+math.Sqrt(discriminant))/2.0*a, sphere),
	}
}
