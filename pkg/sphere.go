package goray

type Sphere struct {
	Center    Point
	Radius    float64
	Transform Matrix
}

func NewSphere() Sphere {
	return Sphere{
		Center:    NewPoint(0, 0, 0),
		Radius:    1.0,
		Transform: IdentityMatrix(),
	}
}

func (s *Sphere) SetTransform(m Matrix) {
	s.Transform = m
}
