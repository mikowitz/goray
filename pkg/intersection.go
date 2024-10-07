package goray

type Intersection struct {
	T      float64
	Object Sphere
}

type Intersections []Intersection

func NewIntersection(t float64, s Sphere) Intersection {
	return Intersection{T: t, Object: s}
}
