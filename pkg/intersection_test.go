package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingAnIntersection(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)

	assert.Equal(t, i.T, 3.5)
	assert.Equal(t, i.Object, s)
}

func TestIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	xs := Intersections{i1, i2}

	assert.Equal(t, len(xs), 2)
	assert.Equal(t, xs[0].T, 1.0)
	assert.Equal(t, xs[1].T, 2.0)
}

func TestHit(t *testing.T) {
	s := NewSphere()

	t.Run("when all intersections have positive t", func(t *testing.T) {
		i1 := NewIntersection(1, s)
		i2 := NewIntersection(2, s)
		xs := Intersections{i2, i1}

		hit, _ := xs.Hit()
		assert.Equal(t, hit, i1)
	})

	t.Run("when some intersections have negative t", func(t *testing.T) {
		i1 := NewIntersection(-1, s)
		i2 := NewIntersection(1, s)
		xs := Intersections{i2, i1}

		hit, _ := xs.Hit()
		assert.Equal(t, hit, i2)
	})

	t.Run("when all intersections have negative t", func(t *testing.T) {
		i1 := NewIntersection(-2, s)
		i2 := NewIntersection(-1, s)
		xs := Intersections{i2, i1}

		_, isHit := xs.Hit()
		assert.False(t, isHit)
	})

	t.Run("the hit is always the lowest nonnegative intersection", func(t *testing.T) {
		i1 := NewIntersection(5, s)
		i2 := NewIntersection(7, s)
		i3 := NewIntersection(-3, s)
		i4 := NewIntersection(2, s)
		xs := Intersections{i1, i2, i3, i4}

		hit, _ := xs.Hit()
		assert.Equal(t, hit, i4)
	})
}
