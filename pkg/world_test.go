package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func defaultWorld() World {
	s1 := NewSphere()
	m := NewMaterial()
	m.Color = NewColor(0.8, 1, 0.6)
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)

	s2 := NewSphere()
	s2.SetTransform(Scaling(0.5, 0.5, 0.5))

	return World{
		LightSource: NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)),
		Objects:     []Sphere{s1, s2},
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

		comps := i.PrepareComputations(r)
		c := w.ShadeHit(comps)

		assert.True(t, TuplesEqual(c, NewColor(0.38066, 0.47583, 0.2855)))
	})

	t.Run("from the inside", func(t *testing.T) {
		w := defaultWorld()
		w.LightSource = NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1))

		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		shape := w.Objects[1]
		i := NewIntersection(0.5, shape)

		comps := i.PrepareComputations(r)
		c := w.ShadeHit(comps)

		assert.True(t, TuplesEqual(c, NewColor(0.90498, 0.90498, 0.90498)))
	})
}

func TestColorAt(t *testing.T) {
	w := defaultWorld()

	t.Run("when a ray misses", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
		c := w.ColorAt(r)
		assert.True(t, TuplesEqual(c, NewColor(0, 0, 0)))
	})

	t.Run("when a ray hits", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		c := w.ColorAt(r)
		assert.True(t, TuplesEqual(c, NewColor(0.38066, 0.47583, 0.2855)))
	})

	t.Run("an intersection behind the ray", func(t *testing.T) {
		w := defaultWorld()
		m := NewMaterial()
		m.Ambient = 1.0
		w.Objects[0].SetMaterial(m)
		w.Objects[1].SetMaterial(m)

		r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
		c := w.ColorAt(r)

		assert.True(t, TuplesEqual(c, w.Objects[1].Material.Color))
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
