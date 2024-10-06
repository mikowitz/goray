package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	assert.True(t, TuplesEqual(c1.Add(c2), NewColor(1.6, 0.7, 1.0)))
}

func TestSubtractingColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	assert.True(t, TuplesEqual(c1.Sub(c2), NewColor(0.2, 0.5, 0.5)))
}

func TestMultiplyingColors(t *testing.T) {
	t.Run("multiplying a color by a scalar", func(t *testing.T) {
		c := NewColor(0.2, 0.3, 0.4)

		assert.True(t, TuplesEqual(c.Mul(2), NewColor(0.4, 0.6, 0.8)))
	})

	t.Run("multiplying colors", func(t *testing.T) {
		c1 := NewColor(1, 0.2, 0.4)
		c2 := NewColor(0.9, 1., 0.1)

		assert.True(t, TuplesEqual(c1.Prod(c2), NewColor(0.9, 0.2, 0.04)))
	})
}
