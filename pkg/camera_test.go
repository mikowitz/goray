package goray

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructingACamera(t *testing.T) {
	c := NewCamera(160, 4.0/3.0, math.Pi/2)

	assert.Equal(t, c.Width, 160)
	assert.Equal(t, c.Height, 120)
	assert.Equal(t, c.FieldOfView, math.Pi/2)
	assert.True(t, MatricesEqual(c.Transform, IdentityMatrix()))
}

func TestPixelSize(t *testing.T) {
	t.Run("for a horizontal camera", func(t *testing.T) {
		c := NewCamera(200, 1.6, math.Pi/2)
		assert.InDelta(t, c.pixelSize, 0.01, 0.00001)
	})

	t.Run("for a vertical camera", func(t *testing.T) {
		c := NewCamera(125, 0.625, math.Pi/2)
		assert.InDelta(t, c.pixelSize, 0.01, 0.00001)
	})
}

func TestRayForPixel(t *testing.T) {
	t.Run("through the center of the canvas", func(t *testing.T) {
		c := NewCamera(201, 201.0/101.0, math.Pi/2)
		r := c.RayForPixel(100, 50)

		assert.True(t, TuplesEqual(r.Origin, NewPoint(0, 0, 0)))
		assert.True(t, TuplesEqual(r.Direction, NewVector(0, 0, -1)))
	})

	t.Run("through a corner of the canvas", func(t *testing.T) {
		c := NewCamera(201, 201.0/101.0, math.Pi/2)
		r := c.RayForPixel(0, 0)

		assert.True(t, TuplesEqual(r.Origin, NewPoint(0, 0, 0)))
		assert.True(t, TuplesEqual(r.Direction, NewVector(0.66519, 0.33259, -0.66851)))
	})

	t.Run("when the camera is transformed", func(t *testing.T) {
		c := NewCamera(201, 201.0/101.0, math.Pi/2)
		c.Transform = RotationY(math.Pi / 4).Mul(Translation(0, -2, 5))
		r := c.RayForPixel(100, 50)

		assert.True(t, TuplesEqual(r.Origin, NewPoint(0, 2, -5)))
		assert.True(t, TuplesEqual(r.Direction, NewVector(math.Sqrt2/2, 0, -math.Sqrt2/2)))
	})
}

func TestRenderingAWorldWithACamera(t *testing.T) {
	w := defaultWorld()
	c := NewCamera(11, 1, math.Pi/2)

	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)

	c.Transform = NewViewTransform(from, to, up)

	image := c.Render(w)

	assert.True(t, TuplesEqual(image.At(5, 5), NewColor(0.38066, 0.47583, 0.2855)))
}
