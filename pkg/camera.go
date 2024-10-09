package goray

import (
	"math"
	"os"

	"github.com/schollz/progressbar/v3"
)

type Camera struct {
	Width, Height         int
	AspectRatio           float64
	FieldOfView           float64
	Transform             Matrix
	halfWidth, halfHeight float64
	pixelSize             float64
}

func NewCamera(width int, aspectRatio, fieldOfView float64) Camera {
	height := int(float64(width) / aspectRatio)
	halfView := math.Tan(fieldOfView / 2.0)
	var halfWidth float64
	var halfHeight float64

	if aspectRatio >= 1.0 {
		halfWidth = halfView
		halfHeight = halfView / aspectRatio
	} else {
		halfWidth = halfView * aspectRatio
		halfHeight = halfView
	}

	return Camera{
		Width:       width,
		Height:      height,
		AspectRatio: aspectRatio,
		FieldOfView: fieldOfView,
		Transform:   IdentityMatrix(),
		halfWidth:   halfWidth,
		halfHeight:  halfHeight,
		pixelSize:   (halfWidth * 2) / float64(width),
	}
}

func (c Camera) RayForPixel(x, y int) Ray {
	xOffset := (float64(x) + 0.5) * c.pixelSize
	yOffset := (float64(y) + 0.5) * c.pixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	pixel := c.Transform.Inverse().Mult(NewPoint(worldX, worldY, -1))
	origin := c.Transform.Inverse().Mult(NewPoint(0, 0, 0))
	direction := pixel.Sub(origin).Normalize()

	return NewRay(origin, direction)
}

func (c Camera) Render(w World) Canvas {
	canvas := NewCanvas(c.Width, c.AspectRatio)

	bar := progressbar.NewOptions(c.Width*c.Height,
		progressbar.OptionSetWriter(os.Stderr),
	)

	for y := range c.Height {
		for x := range c.Width {
			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray, 5)
			canvas.Write(x, y, color)
			err := bar.Add(1)
			if err != nil {
				panic(err)
			}
		}
	}
	return canvas
}
