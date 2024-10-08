package goray

import "math"

type Sphere struct {
	Center    Point
	Radius    float64
	Transform Matrix
	Material  Material
	SavedRay  Ray
}

func NewSphere() Sphere {
	return Sphere{
		Center:    NewPoint(0, 0, 0),
		Radius:    1.0,
		Transform: IdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func (s *Sphere) GetTransform() Matrix {
	return s.Transform
}

func (s *Sphere) SetTransform(m Matrix) {
	s.Transform = m
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m Material) {
	s.Material = m
}

func (s *Sphere) GetSavedRay() Ray {
	return s.SavedRay
}

func (s *Sphere) LocalIntersect(r Ray) Intersections {
	s.SavedRay = r
	sphereToRay := r.Origin.Sub(NewPoint(0, 0, 0))

	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := (b * b) - (4.0 * a * c)

	if discriminant < 0.0 {
		return Intersections{}
	}

	return Intersections{
		NewIntersection((-b-math.Sqrt(discriminant))/(2.0*a), s),
		NewIntersection((-b+math.Sqrt(discriminant))/(2.0*a), s),
	}
}

func (s *Sphere) LocalNormalAt(p Point) Vector {
	return p.Sub(NewPoint(0, 0, 0))
}
