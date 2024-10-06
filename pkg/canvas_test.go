package goray

import (
	"strings"
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

func TestConstructingPPMPixelData(t *testing.T) {
	c := NewCanvas(5, 5.0/3.0)
	c1 := NewColor(1.5, 0, 0)
	c2 := NewColor(0, 0.5, 0)
	c3 := NewColor(-0.5, 0, 1)

	c.Write(0, 0, c1)
	c.Write(2, 1, c2)
	c.Write(4, 2, c3)

	ppm := c.ToPpm()

	lines := strings.Split(ppm, "\n")

	assert.Equal(t, lines[0], "PPM")
	assert.Equal(t, lines[1], "5 3")
	assert.Equal(t, lines[2], "255")

	for i := 3; i < 18; i++ {
		if i == 3 {
			assert.Equal(t, lines[i], "255 0 0")
		} else if i == 10 {
			assert.Equal(t, lines[i], "0 128 0")
		} else if i == 17 {
			assert.Equal(t, lines[i], "0 0 255")
		} else {
			assert.Equal(t, lines[i], "0 0 0")
		}
	}
}
