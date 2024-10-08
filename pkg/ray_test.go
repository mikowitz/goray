package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayAt(t *testing.T) {
	r := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))

	assert.True(t, TuplesEqual(r.At(0), NewPoint(2, 3, 4)))
	assert.True(t, TuplesEqual(r.At(1), NewPoint(3, 3, 4)))
	assert.True(t, TuplesEqual(r.At(-1), NewPoint(1, 3, 4)))
	assert.True(t, TuplesEqual(r.At(2.5), NewPoint(4.5, 3, 4)))
}

func TestIntersectingRaysWithSpheres(t *testing.T) {
	t.Run("a ray intersects a sphere at two points", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 2)
		assert.Equal(t, xs[0].T, 4.0)
		assert.Equal(t, xs[0].Object, &s)
		assert.Equal(t, xs[1].T, 6.0)
		assert.Equal(t, xs[1].Object, &s)
	})

	t.Run("a ray intersects a sphere at a tangent", func(t *testing.T) {
		r := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 2)
		assert.Equal(t, xs[0].T, 5.0)
		assert.Equal(t, xs[1].T, 5.0)
	})

	t.Run("a ray misses a sphere", func(t *testing.T) {
		r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 0)
	})

	t.Run("a ray originates inside a sphere", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 2)
		assert.Equal(t, xs[0].T, -1.0)
		assert.Equal(t, xs[1].T, 1.0)
	})

	t.Run("a ray originates in front of a sphere", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
		s := NewSphere()

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 2)
		assert.Equal(t, xs[0].T, -6.0)
		assert.Equal(t, xs[1].T, -4.0)
	})

	t.Run("intersecting a scaled sphere with a ray", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()
		s.SetTransform(Scaling(2, 2, 2))

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 2)
		assert.Equal(t, xs[0].T, 3.0)
		assert.Equal(t, xs[1].T, 7.0)
	})

	t.Run("intersecting a translated sphere with a ray", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewSphere()
		s.SetTransform(Translation(5, 0, 0))

		xs := r.Intersect(&s)

		assert.Equal(t, len(xs), 0)
	})
}

func TestTransformingARay(t *testing.T) {
	t.Run("translating a ray", func(t *testing.T) {
		r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
		m := Translation(3, 4, 5)

		r2 := r.Transform(m)

		assert.True(t, TuplesEqual(r2.Origin, NewPoint(4, 6, 8)))
		assert.True(t, TuplesEqual(r2.Direction, NewVector(0, 1, 0)))
	})

	t.Run("scaling a ray", func(t *testing.T) {
		r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
		m := Scaling(2, 3, 4)

		r2 := r.Transform(m)

		assert.True(t, TuplesEqual(r2.Origin, NewPoint(2, 6, 12)))
		assert.True(t, TuplesEqual(r2.Direction, NewVector(0, 3, 0)))
	})
}
