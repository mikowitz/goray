package goray

import (
	"cmp"
	"slices"
)

type Intersection struct {
	T      float64
	Object Sphere
}

type Intersections []Intersection

type Computations struct {
	Object        Sphere
	T             float64
	Point         Point
	OverPoint     Point
	Eyev, Normalv Vector
	Inside        bool
}

func NewIntersection(t float64, s Sphere) Intersection {
	return Intersection{T: t, Object: s}
}

func (i Intersection) PrepareComputations(ray Ray) Computations {
	point := ray.At(i.T)
	eyev := ray.Direction.Neg()
	normalv := i.Object.NormalAt(point)
	inside := false

	if normalv.Dot(eyev) < 0 {
		inside = true
		normalv = normalv.Neg()
	}

	return Computations{
		Object:    i.Object,
		T:         i.T,
		Point:     point,
		OverPoint: point.Add(normalv.Mul(0.00001)),
		Eyev:      eyev,
		Normalv:   normalv,
		Inside:    inside,
	}
}

func (xs Intersections) Hit() (Intersection, bool) {
	slices.SortFunc(xs, func(a, b Intersection) int {
		return cmp.Compare(a.T, b.T)
	})
	i := slices.IndexFunc(xs, func(x Intersection) bool {
		return x.T >= 0.0
	})
	if i == -1 {
		return Intersection{}, false
	}
	return xs[i], true
}
