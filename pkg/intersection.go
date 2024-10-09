package goray

import (
	"cmp"
	"math"
	"slices"
)

type Intersection struct {
	T      float64
	Object Shape
}

type Intersections []Intersection

type Computations struct {
	Object                  Shape
	T                       float64
	Point                   Point
	OverPoint, UnderPoint   Point
	Eyev, Normalv, Reflectv Vector
	Inside                  bool
	N1, N2                  float64
}

func NewIntersection(t float64, s Shape) Intersection {
	return Intersection{T: t, Object: s}
}

func (i Intersection) PrepareComputations(ray Ray, xs Intersections) Computations {
	point := ray.At(i.T)
	eyev := ray.Direction.Neg()
	normalv := NormalAt(i.Object, point)
	reflectv := ray.Direction.Reflect(normalv)
	inside := false

	if normalv.Dot(eyev) < 0 {
		inside = true
		normalv = normalv.Neg()
	}

	var n1 float64
	var n2 float64

	containers := []Shape{}
	for _, x := range xs {
		if x == i {
			if len(containers) == 0 {
				n1 = 1.0
			} else {
				n1 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}

		if slices.Contains(containers, x.Object) {
			containers = slices.DeleteFunc(containers, func(s Shape) bool {
				return s == x.Object
			})
		} else {
			containers = append(containers, x.Object)
		}

		if x == i {
			if len(containers) == 0 {
				n2 = 1.0
			} else {
				n2 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}
	}

	return Computations{
		Object:     i.Object,
		T:          i.T,
		Point:      point,
		OverPoint:  point.Add(normalv.Mul(0.00001)),
		UnderPoint: point.Sub(normalv.Mul(0.00001)),
		Eyev:       eyev,
		Normalv:    normalv,
		Reflectv:   reflectv,
		Inside:     inside,
		N1:         n1,
		N2:         n2,
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

func (c Computations) Schlick() float64 {
	cos := c.Eyev.Dot(c.Normalv)

	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sin2T := n * n * (1.0 - cos*cos)
		if sin2T > 1.0 {
			return 1.0
		}

		cosT := math.Sqrt(1.0 - sin2T)
		cos = cosT
	}
	r0 := math.Pow((c.N1-c.N2)/(c.N1+c.N2), 2)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
