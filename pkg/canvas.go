package goray

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
