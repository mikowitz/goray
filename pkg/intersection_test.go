package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingAnIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, &s)

	assert.Equal(t, i.T, 3.5)
	assert.Equal(t, i.Object, &s)
}

func TestIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, &s)
	i2 := NewIntersection(2, &s)

	xs := Intersections{i1, i2}

	assert.Equal(t, len(xs), 2)
	assert.Equal(t, xs[0].T, 1.0)
	assert.Equal(t, xs[1].T, 2.0)
}

func TestHit(t *testing.T) {
	s := NewSphere()

	t.Run("when all intersections have positive t", func(t *testing.T) {
		i1 := NewIntersection(1, &s)
		i2 := NewIntersection(2, &s)
		xs := Intersections{i2, i1}

		hit, _ := xs.Hit()
		assert.Equal(t, hit, i1)
	})

	t.Run("when some intersections have negative t", func(t *testing.T) {
		i1 := NewIntersection(-1, &s)
		i2 := NewIntersection(1, &s)
		xs := Intersections{i2, i1}

		hit, _ := xs.Hit()
		assert.Equal(t, hit, i2)
	})

	t.Run("when all intersections have negative t", func(t *testing.T) {
		i1 := NewIntersection(-2, &s)
		i2 := NewIntersection(-1, &s)
		xs := Intersections{i2, i1}

		_, isHit := xs.Hit()
		assert.False(t, isHit)
	})

	t.Run("the hit is always the lowest nonnegative intersection", func(t *testing.T) {
		i1 := NewIntersection(5, &s)
		i2 := NewIntersection(7, &s)
		i3 := NewIntersection(-3, &s)
		i4 := NewIntersection(2, &s)
		xs := Intersections{i1, i2, i3, i4}

		hit, _ := xs.Hit()
		assert.Equal(t, hit, i4)
	})
}

func TestPrecomputingIntersectionState(t *testing.T) {
	t.Run("when an intersection occurs on the outside", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		shape := NewSphere()
		i := NewIntersection(4, &shape)

		comps := i.PrepareComputations(r, Intersections{i})

		assert.Equal(t, comps.T, i.T)
		assert.True(t, TuplesEqual(comps.Point, NewPoint(0, 0, -1)))
		assert.True(t, TuplesEqual(comps.Eyev, NewVector(0, 0, -1)))
		assert.True(t, TuplesEqual(comps.Normalv, NewVector(0, 0, -1)))
		assert.False(t, comps.Inside)
	})

	t.Run("when an intersection occurs on the inside", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		shape := NewSphere()
		i := NewIntersection(1, &shape)

		comps := i.PrepareComputations(r, Intersections{i})

		assert.Equal(t, comps.T, i.T)
		assert.True(t, TuplesEqual(comps.Point, NewPoint(0, 0, 1)))
		assert.True(t, TuplesEqual(comps.Eyev, NewVector(0, 0, -1)))
		assert.True(t, TuplesEqual(comps.Normalv, NewVector(0, 0, -1)))
		assert.True(t, comps.Inside)
	})

	t.Run("the hit should offset the point", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		shape := NewSphere()
		shape.SetTransform(Translation(0, 0, 1))
		i := NewIntersection(5, &shape)

		comps := i.PrepareComputations(r, Intersections{i})
		assert.True(t, comps.OverPoint.z < -0.000005)
		assert.True(t, comps.Point.z > comps.OverPoint.z)

	})

	t.Run("the underpoint is offset below the surface", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		shape := GlassSphere()
		shape.SetTransform(Translation(0, 0, 1))
		i := NewIntersection(5, &shape)

		comps := i.PrepareComputations(r, Intersections{i})
		assert.True(t, comps.UnderPoint.z > 0.000005)
		assert.True(t, comps.Point.z < comps.UnderPoint.z)

	})

	t.Run("the reflection vector", func(t *testing.T) {
		s := NewPlane()
		r := NewRay(NewPoint(0, 1, -1), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
		i := NewIntersection(math.Sqrt2/2, &s)

		comps := i.PrepareComputations(r, Intersections{i})

		assert.True(t, TuplesEqual(comps.Reflectv, NewVector(0, math.Sqrt2/2, math.Sqrt2/2)))
	})
}

func TestFindingN1AndN2(t *testing.T) {
	a := GlassSphere()
	a.SetTransform(Scaling(2, 2, 2))
	a.Material.RefractiveIndex = 1.5

	b := GlassSphere()
	b.SetTransform(Translation(0, 0, -0.25))
	b.Material.RefractiveIndex = 2.0

	c := GlassSphere()
	c.SetTransform(Translation(0, 0, 0.25))
	c.Material.RefractiveIndex = 2.5

	r := NewRay(NewPoint(0, 0, -4), NewVector(0, 0, 1))
	xs := Intersections{
		NewIntersection(2, &a),
		NewIntersection(2.75, &b),
		NewIntersection(3.25, &c),
		NewIntersection(4.75, &b),
		NewIntersection(5.25, &c),
		NewIntersection(6, &a),
	}

	testCases := [][2]float64{
		[2]float64{1, 1.5},
		[2]float64{1.5, 2},
		[2]float64{2, 2.5},
		[2]float64{2.5, 2.5},
		[2]float64{2.5, 1.5},
		[2]float64{1.5, 1.0},
	}

	for i, tc := range testCases {
		comps := xs[i].PrepareComputations(r, xs)
		assert.Equal(t, comps.N1, tc[0])
		assert.Equal(t, comps.N2, tc[1])
	}
}
