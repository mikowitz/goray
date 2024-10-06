package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingACanvas(t *testing.T) {
	c := NewCanvas(10, 0.5)

	assert.Equal(t, c.Width, 10)
	assert.Equal(t, c.Height, 20)
	assert.Equal(t, len(c.Pixels), 200)

	for _, p := range c.Pixels {
		assert.True(t, TuplesEqual(p, NewColor(0, 0, 0)))
	}
}

func TestWritingPixelsToACanvas(t *testing.T) {
	c := NewCanvas(10, 0.5)
	red := NewColor(1, 0, 0)

	c.Write(2, 3, red)
	assert.True(t, TuplesEqual(c.At(2, 3), red))
}
