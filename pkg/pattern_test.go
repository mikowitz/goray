package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripePattern(t *testing.T) {
	pattern := NewStripePattern(White(), Black())

	t.Run("is constant in y", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 1, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 2, 0)), White())
	})

	t.Run("is constant in z", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 0, 1)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 0, 2)), White())
	})

	t.Run("is constant in x", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0.9, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(1, 0, 0)), Black())
		assert.Equal(t, pattern.At(NewPoint(-0.1, 0, 0)), Black())
		assert.Equal(t, pattern.At(NewPoint(-1, 0, 0)), Black())
		assert.Equal(t, pattern.At(NewPoint(-1.1, 0, 0)), White())
	})

	t.Run("with an object transformation", func(t *testing.T) {
		s := NewSphere()
		s.SetTransform(Scaling(2, 2, 2))
		c := pattern.AtObject(&s, NewPoint(1.5, 0, 0))

		assert.Equal(t, c, White())
	})

	t.Run("with a pattern transformation", func(t *testing.T) {
		pattern.SetTransform(Scaling(2, 2, 2))
		s := NewSphere()
		c := pattern.AtObject(&s, NewPoint(1.5, 0, 0))

		assert.Equal(t, c, White())
	})

	t.Run("with both an object and pattern transform", func(t *testing.T) {
		s := NewSphere()
		s.SetTransform(Scaling(2, 2, 2))
		pattern.SetTransform(Translation(0.5, 0, 0))
		c := pattern.AtObject(&s, NewPoint(2.5, 0, 0))

		assert.Equal(t, c, White())
	})

}
