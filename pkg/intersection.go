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

func NewIntersection(t float64, s Sphere) Intersection {
	return Intersection{T: t, Object: s}
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
