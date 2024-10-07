package goray

import (
	"fmt"
	"strings"
)

type Canvas struct {
	Width, Height int
	Pixels        []Color
}

func NewCanvas(width int, aspectRatio float64) Canvas {
	height := int(float64(width) / aspectRatio)
	return Canvas{
		Width:  width,
		Height: height,
		Pixels: make([]Tuple, width*height),
	}
}

func (c Canvas) Write(x, y int, color Color) {
	c.Pixels[y*c.Width+x] = color
}

func (c Canvas) At(x, y int) Color {
	return c.Pixels[y*c.Width+x]
}

func (c Canvas) ToPpm() string {
	pixels := make([]string, c.Height*c.Width)

	for i, p := range c.Pixels {
		pixels[i] = p.ToPpm()
	}

	return fmt.Sprintf("P3\n%d %d\n255\n%s\n", c.Width, c.Height, strings.Join(pixels, "\n"))
}
