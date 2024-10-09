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
		c := PatternAtObject(&pattern, &s, NewPoint(1.5, 0, 0))

		assert.Equal(t, c, White())
	})

	t.Run("with a pattern transformation", func(t *testing.T) {
		pattern.SetTransform(Scaling(2, 2, 2))
		s := NewSphere()
		c := PatternAtObject(&pattern, &s, NewPoint(1.5, 0, 0))

		assert.Equal(t, c, White())
	})

	t.Run("with both an object and pattern transform", func(t *testing.T) {
		s := NewSphere()
		s.SetTransform(Scaling(2, 2, 2))
		pattern.SetTransform(Translation(0.5, 0, 0))
		c := PatternAtObject(&pattern, &s, NewPoint(2.5, 0, 0))

		assert.Equal(t, c, White())
	})
}

func TestGradientPattern(t *testing.T) {
	pattern := NewGradientPattern(White(), Black())
	assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
	assert.Equal(t, pattern.At(NewPoint(0.25, 0, 0)), NewColor(0.75, 0.75, 0.75))
	assert.Equal(t, pattern.At(NewPoint(0.5, 0, 0)), NewColor(0.5, 0.5, 0.5))
	assert.Equal(t, pattern.At(NewPoint(0.75, 0, 0)), NewColor(0.25, 0.25, 0.25))
}

func TestRingPattern(t *testing.T) {
	pattern := NewRingPattern(White(), Black())
	assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
	assert.Equal(t, pattern.At(NewPoint(1, 0, 0)), Black())
	assert.Equal(t, pattern.At(NewPoint(0, 0, 1)), Black())
	assert.Equal(t, pattern.At(NewPoint(0.708, 0, 0.708)), Black())
}

func TestCheckersPattern(t *testing.T) {
	pattern := NewCheckersPattern(White(), Black())

	t.Run("repeats in x", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0.99, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(1.01, 0, 0)), Black())
	})

	t.Run("repeats in y", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 0.99, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 1.01, 0)), Black())
	})

	t.Run("repeats in z", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 0, 0.99)), White())
		assert.Equal(t, pattern.At(NewPoint(0, 0, 1.01)), Black())
	})

	t.Run("repeats in all directions", func(t *testing.T) {
		assert.Equal(t, pattern.At(NewPoint(0.99, 0, 0.99)), White())
		assert.Equal(t, pattern.At(NewPoint(1.01, 0, 0.99)), Black())
		assert.Equal(t, pattern.At(NewPoint(1.01, 0, 1.01)), White())
	})
}
