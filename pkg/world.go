package goray

import (
	"cmp"
	"slices"
)

type World struct {
	LightSource PointLight
	Objects     []Sphere
}

func NewWorld() World {
	return World{}
}

func (w World) Intersect(ray Ray) Intersections {
	xs := Intersections{}
	for _, object := range w.Objects {
		xs = append(xs, ray.Intersect(object)...)
	}
	slices.SortFunc(xs, func(a, b Intersection) int {
		return cmp.Compare(a.T, b.T)
	})
	return xs
}

func (w World) ShadeHit(c Computations) Color {
	return c.Object.Material.Lighting(w.LightSource, c.Point, c.Eyev, c.Normalv)
}

func (w World) ColorAt(r Ray) Color {
	xs := w.Intersect(r)
	if hit, isHit := xs.Hit(); isHit {
		comps := hit.PrepareComputations(r)
		return w.ShadeHit(comps)
	}
	return NewColor(0, 0, 0)
}
