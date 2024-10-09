package goray

import (
	"cmp"
	"math"
	"slices"
)

type World struct {
	LightSource PointLight
	Objects     []Shape
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

func (w World) ShadeHit(c Computations, depth int) Color {
	inShadow := w.IsShadowed(c.OverPoint)
	surface := c.Object.GetMaterial().Lighting(c.Object, w.LightSource, c.Point, c.Eyev, c.Normalv, inShadow)

	reflected := w.ReflectedColor(c, depth)
	refracted := w.RefractedColor(c, depth)

	material := c.Object.GetMaterial()
	if material.Reflective > 0.0 && material.Transparency > 0.0 {
		reflectance := c.Schlick()
		return surface.Add(reflected.Mul(reflectance)).Add(refracted.Mul(1.0 - reflectance))
	} else {
		return surface.Add(reflected).Add(refracted)
	}
}

func (w World) ColorAt(r Ray, depth int) Color {
	xs := w.Intersect(r)
	if hit, isHit := xs.Hit(); isHit {
		comps := hit.PrepareComputations(r, xs)
		return w.ShadeHit(comps, depth)
	}
	return NewColor(0, 0, 0)
}

func (w World) IsShadowed(point Point) bool {
	v := w.LightSource.Position.Sub(point)
	distance := v.Magnitude()
	ray := NewRay(point, v.Normalize())
	xs := w.Intersect(ray)

	if hit, isHit := xs.Hit(); isHit {
		return hit.T < distance
	}
	return false
}

func (w World) ReflectedColor(c Computations, depth int) Color {
	if depth <= 0 {
		return Black()
	}
	if c.Object.GetMaterial().Reflective == 0 {
		return Black()
	}
	reflectRay := NewRay(c.OverPoint, c.Reflectv)
	color := w.ColorAt(reflectRay, depth-1)

	return color.Mul(c.Object.GetMaterial().Reflective)
}

func (w World) RefractedColor(c Computations, depth int) Color {
	if depth <= 0 {
		return Black()
	}
	if c.Object.GetMaterial().Transparency == 0.0 {
		return Black()
	}
	nRatio := c.N1 / c.N2
	cosI := c.Eyev.Dot(c.Normalv)

	sin2T := nRatio * nRatio * (1 - cosI*cosI)

	if sin2T > 1.0 {
		return Black()
	}

	cosT := math.Sqrt(1.0 - sin2T)
	direction := c.Normalv.Mul(nRatio*cosI - cosT).Sub(c.Eyev.Mul(nRatio))

	refractRay := NewRay(c.UnderPoint, direction)
	return w.ColorAt(refractRay, depth-1).Mul(c.Object.GetMaterial().Transparency)
}
