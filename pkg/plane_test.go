package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaneNormal(t *testing.T) {
	p := NewPlane()
	n1 := p.LocalNormalAt(NewPoint(0, 0, 0))
	n2 := p.LocalNormalAt(NewPoint(10, 0, -10))
	n3 := p.LocalNormalAt(NewPoint(-5, 0, 150))

	assert.True(t, TuplesEqual(n1, NewVector(0, 1, 0)))
	assert.True(t, TuplesEqual(n2, NewVector(0, 1, 0)))
	assert.True(t, TuplesEqual(n3, NewVector(0, 1, 0)))
}

func TestPlaneIntersection(t *testing.T) {
	t.Run("with a parallel ray", func(t *testing.T) {
		p := NewPlane()
		r := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))
		xs := p.LocalIntersect(r)

		assert.Empty(t, xs)
	})

	t.Run("with a coplanar ray", func(t *testing.T) {
		p := NewPlane()
		r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
		xs := p.LocalIntersect(r)

		assert.Empty(t, xs)
	})

	t.Run("from above", func(t *testing.T) {
		p := NewPlane()
		r := NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0))
		xs := p.LocalIntersect(r)

		assert.Equal(t, len(xs), 1)
		assert.Equal(t, xs[0].T, 1.0)
		assert.Equal(t, xs[0].Object, &p)
	})

	t.Run("from below", func(t *testing.T) {
		p := NewPlane()
		r := NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0))
		xs := p.LocalIntersect(r)

		assert.Equal(t, len(xs), 1)
		assert.Equal(t, xs[0].T, 1.0)
		assert.Equal(t, xs[0].Object, &p)
	})
}
