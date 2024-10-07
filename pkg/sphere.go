package goray

type Sphere struct {
	Center    Point
	Radius    float64
	Transform Matrix
	Material  Material
}

func NewSphere() Sphere {
	return Sphere{
		Center:    NewPoint(0, 0, 0),
		Radius:    1.0,
		Transform: IdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func (s *Sphere) SetTransform(m Matrix) {
	s.Transform = m
}

func (s *Sphere) SetMaterial(m Material) {
	s.Material = m
}

func (s Sphere) NormalAt(p Point) Vector {
	objectPoint := s.Transform.Inverse().Mult(p)
	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := s.Transform.Inverse().Transpose().Mult(objectNormal)
	worldNormal.w = 0.0
	return worldNormal.Normalize()
}
