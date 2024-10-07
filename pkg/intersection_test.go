package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingAnIntersection(t *testing.T) {
	s := DemoSphere()
	i := NewIntersection(3.5, s)

	assert.Equal(t, i.T, 3.5)
	assert.Equal(t, i.Object, s)
}

func TestIntersections(t *testing.T) {
	s := DemoSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	xs := Intersections{i1, i2}

	assert.Equal(t, len(xs), 2)
	assert.Equal(t, xs[0].T, 1.0)
	assert.Equal(t, xs[1].T, 2.0)
}
