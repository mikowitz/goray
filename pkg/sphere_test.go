package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSphereTransform(t *testing.T) {
	t.Run("a sphere's default transformation", func(t *testing.T) {
		s := NewSphere()
		assert.True(t, MatricesEqual(s.Transform, IdentityMatrix()))
	})

	t.Run("changing a sphere's transformation", func(t *testing.T) {
		s := NewSphere()
		transform := Translation(2, 3, 4)
		s.SetTransform(transform)

		assert.True(t, MatricesEqual(s.Transform, transform))
	})
}
