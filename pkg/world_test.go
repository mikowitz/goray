package goray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultWorld() World {
	s1 := NewSphere()
	pattern := NewSolidPattern(NewColor(0.8, 1, 0.6))
	m := NewMaterial()
	m.Pattern = &pattern
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)

	s2 := NewSphere()
	s2.SetTransform(Scaling(0.5, 0.5, 0.5))

	return World{
		LightSource: NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		Objects:     []Shape{&s1, &s2},
	}
}

func TestEmptyWorld(t *testing.T) {
	w := NewWorld()

	assert.Empty(t, w.Objects)
}

func TestIntersectWorldWithRay(t *testing.T) {
	w := defaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	xs := w.Intersect(r)

	assert.Equal(t, len(xs), 4)
	assert.Equal(t, xs[0].T, 4.0)
	assert.Equal(t, xs[1].T, 4.5)
	assert.Equal(t, xs[2].T, 5.5)
	assert.Equal(t, xs[3].T, 6.0)
}

func TestShadeHit(t *testing.T) {
	t.Run("from the outside", func(t *testing.T) {
		w := defaultWorld()
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		shape := w.Objects[0]
		i := NewIntersection(4, shape)

		comps := i.PrepareComputations(r, Intersections{i})
		c := w.ShadeHit(comps, 0)

		assert.True(t, TuplesEqual(c, NewColor(0.38066, 0.47583, 0.2855)))
	})

	t.Run("from the inside", func(t *testing.T) {
		w := defaultWorld()
		w.LightSource = NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1))

		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		shape := w.Objects[1]
		i := NewIntersection(0.5, shape)

		comps := i.PrepareComputations(r, Intersections{i})
		c := w.ShadeHit(comps, 0)

		assert.True(t, TuplesEqual(c, NewColor(0.90498, 0.90498, 0.90498)))
	})

	t.Run("for an intersection in shadow", func(t *testing.T) {
		w := NewWorld()
		w.LightSource = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
		s1 := NewSphere()
		s2 := NewSphere()
		s2.SetTransform(Translation(0, 0, 10))
		w.Objects = []Shape{&s1, &s2}

		r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
		i := NewIntersection(4, &s2)
		comps := i.PrepareComputations(r, Intersections{i})
		c := w.ShadeHit(comps, 0)

		assert.True(t, TuplesEqual(c, NewColor(0.1, 0.1, 0.1)))
	})

	t.Run("with a reflective material", func(t *testing.T) {
		w := defaultWorld()
		s := NewPlane()
		s.Material.Reflective = 0.5
		s.SetTransform(Translation(0, -1, 0))
		w.Objects = append(w.Objects, &s)

		r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
		i := NewIntersection(math.Sqrt2, &s)

		comps := i.PrepareComputations(r, Intersections{i})

		c := w.ShadeHit(comps, 1)

		assert.True(t, TuplesEqual(c, NewColor(0.87675, 0.92434, 0.82917)))
	})

	t.Run("with a transparent material", func(t *testing.T) {
		w := defaultWorld()
		floor := NewPlane()
		floor.SetTransform(Translation(0, -1, 0))
		floor.Material.Transparency = 0.5
		floor.Material.RefractiveIndex = 1.5

		ball := NewSphere()
		red := NewSolidPattern(NewColor(1, 0, 0))
		ball.Material.Pattern = &red
		ball.Material.Ambient = 0.5
		ball.SetTransform(Translation(0, -3.5, -0.5))

		w.Objects = append(w.Objects, &floor, &ball)

		r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
		xs := Intersections{
			NewIntersection(math.Sqrt2, &floor),
		}

		comps := xs[0].PrepareComputations(r, xs)
		c := w.ShadeHit(comps, 5)
		assert.True(t, TuplesEqual(c, NewColor(0.93642, 0.68642, 0.68642)))
	})

	t.Run("with a reflective, transparent material", func(t *testing.T) {
		w := defaultWorld()
		floor := NewPlane()
		floor.SetTransform(Translation(0, -1, 0))
		floor.Material.Reflective = 0.5
		floor.Material.Transparency = 0.5
		floor.Material.RefractiveIndex = 1.5

		ball := NewSphere()
		red := NewSolidPattern(NewColor(1, 0, 0))
		ball.Material.Pattern = &red
		ball.Material.Ambient = 0.5
		ball.SetTransform(Translation(0, -3.5, -0.5))

		w.Objects = append(w.Objects, &floor, &ball)

		r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
		xs := Intersections{
			NewIntersection(math.Sqrt2, &floor),
		}

		comps := xs[0].PrepareComputations(r, xs)
		c := w.ShadeHit(comps, 5)
		assert.True(t, TuplesEqual(c, NewColor(0.93391, 0.69643, 0.69243)))
	})
}

func TestColorAt(t *testing.T) {
	w := defaultWorld()

	t.Run("when a ray misses", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
		c := w.ColorAt(r, 0)
		assert.True(t, TuplesEqual(c, NewColor(0, 0, 0)))
	})

	t.Run("when a ray hits", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		c := w.ColorAt(r, 0)
		assert.True(t, TuplesEqual(c, NewColor(0.38066, 0.47583, 0.2855)))
	})

	t.Run("an intersection behind the ray", func(t *testing.T) {
		w := defaultWorld()
		m := NewMaterial()
		m.Ambient = 1.0
		w.Objects[0].SetMaterial(m)
		w.Objects[1].SetMaterial(m)

		r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
		c := w.ColorAt(r, 0)

		assert.True(t, TuplesEqual(c, w.Objects[1].GetMaterial().Pattern.At(r.Origin)))
	})
}

func TestIsShadowed(t *testing.T) {
	w := defaultWorld()
	testCases := map[Point]bool{
		NewPoint(0, 10, 0):     false,
		NewPoint(10, -10, 10):  true,
		NewPoint(-20, 20, -20): false,
		NewPoint(-2, 2, -2):    false,
	}

	for p, b := range testCases {
		if b {
			assert.True(t, w.IsShadowed(p))
		} else {
			assert.False(t, w.IsShadowed(p))
		}
	}

}

func TestReflectedColor(t *testing.T) {
	t.Run("for a nonreflective material", func(t *testing.T) {
		w := defaultWorld()
		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		m := w.Objects[1].GetMaterial()
		m.Ambient = 1.0
		w.Objects[1].SetMaterial(m)

		i := NewIntersection(1, w.Objects[1])
		comps := i.PrepareComputations(r, Intersections{i})

		c := w.ReflectedColor(comps, 0)

		assert.Equal(t, c, NewColor(0, 0, 0))
	})

	t.Run("for a reflective material", func(t *testing.T) {
		w := defaultWorld()
		s := NewPlane()
		s.Material.Reflective = 0.5
		s.SetTransform(Translation(0, -1, 0))
		w.Objects = append(w.Objects, &s)

		r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
		i := NewIntersection(math.Sqrt2, &s)

		comps := i.PrepareComputations(r, Intersections{i})

		c := w.ReflectedColor(comps, 1)

		assert.True(t, TuplesEqual(c, NewColor(0.19033, 0.23791, 0.14274)))
	})

	t.Run("at maximum recursive depth", func(t *testing.T) {
		w := defaultWorld()
		s := NewPlane()
		s.Material.Reflective = 0.5
		s.SetTransform(Translation(0, -1, 0))
		w.Objects = append(w.Objects, &s)

		r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
		i := NewIntersection(math.Sqrt2, &s)

		comps := i.PrepareComputations(r, Intersections{i})

		c := w.ReflectedColor(comps, 0)

		assert.True(t, TuplesEqual(c, Black()))
	})
}

func TestRefractedColor(t *testing.T) {
	t.Run("with an opaque surface", func(t *testing.T) {
		w := defaultWorld()
		shape := w.Objects[0]

		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

		xs := Intersections{
			NewIntersection(4, shape),
			NewIntersection(6, shape),
		}

		comps := xs[0].PrepareComputations(r, xs)

		c := w.RefractedColor(comps, 5)

		assert.Equal(t, c, Black())
	})

	t.Run("at maximum recursive depth", func(t *testing.T) {
		w := defaultWorld()
		m := w.Objects[0].GetMaterial()
		m.Transparency = 1.0
		m.RefractiveIndex = 1.5
		w.Objects[0].SetMaterial(m)
		shape := w.Objects[0]

		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

		xs := Intersections{
			NewIntersection(4, shape),
			NewIntersection(6, shape),
		}

		comps := xs[0].PrepareComputations(r, xs)

		c := w.RefractedColor(comps, 0)

		assert.Equal(t, c, Black())
	})

	t.Run("under total internal reflection", func(t *testing.T) {
		w := defaultWorld()
		m := w.Objects[0].GetMaterial()
		m.Transparency = 1.0
		m.RefractiveIndex = 1.5
		w.Objects[0].SetMaterial(m)
		shape := w.Objects[0]

		r := NewRay(NewPoint(0, 0, math.Sqrt2/2), NewVector(0, 1, 0))

		xs := Intersections{
			NewIntersection(-math.Sqrt2/2, shape),
			NewIntersection(math.Sqrt2/2, shape),
		}

		comps := xs[1].PrepareComputations(r, xs)
		c := w.ReflectedColor(comps, 5)

		assert.Equal(t, c, Black())
	})

	t.Run("with a refracted ray", func(t *testing.T) {
		w := defaultWorld()
		m := w.Objects[0].GetMaterial()
		m.Ambient = 1.0
		pattern := DemoPattern{Transform: IdentityMatrix()}
		m.Pattern = &pattern
		w.Objects[0].SetMaterial(m)

		m = w.Objects[1].GetMaterial()
		m.Transparency = 1.0
		m.RefractiveIndex = 1.5
		w.Objects[1].SetMaterial(m)

		r := NewRay(NewPoint(0, 0, 0.1), NewVector(0, 1, 0))

		xs := Intersections{
			NewIntersection(-0.9899, w.Objects[0]),
			NewIntersection(-0.4899, w.Objects[1]),
			NewIntersection(0.4899, w.Objects[1]),
			NewIntersection(0.9899, w.Objects[0]),
		}

		comps := xs[2].PrepareComputations(r, xs)
		c := w.RefractedColor(comps, 5)
		fmt.Println(c)
		assert.True(t, TuplesEqual(c, NewColor(0, 0.99888, 0.04722)))
	})
}

type DemoPattern struct {
	Transform Matrix
}

func (dp *DemoPattern) At(p Point) Color {
	return NewColor(p.x, p.y, p.z)
}

func (dp *DemoPattern) AtObject(shape Shape, point Point) Color {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	patternPoint := dp.GetTransform().Inverse().Mult(objectPoint)
	return dp.At(patternPoint)
}

func (dp *DemoPattern) GetTransform() Matrix {
	return dp.Transform
}

func (dp *DemoPattern) SetTransform(m Matrix) {
	dp.Transform = m
}
