package main

import (
	"fmt"

	g "github.com/mikowitz/goray/pkg"
)

func main() {

	rayOrigin := g.NewPoint(0, 0, -5)

	wallZ := 10.0

	wallSize := 7.0

	canvasPixels := 400

	pixelSize := wallSize / float64(canvasPixels)

	half := wallSize / 2.0

	canvas := g.NewCanvas(canvasPixels, 1.0)
	color := g.NewColor(1, 0, 0)
	shape := g.NewSphere()

	for y := range canvas.Height {
		worldY := half - pixelSize*float64(y)

		for x := range canvas.Width {
			worldX := -half + pixelSize*float64(x)

			position := g.NewPoint(worldX, worldY, wallZ)

			r := g.NewRay(rayOrigin, position.Sub(rayOrigin).Normalize())
			xs := r.Intersect(shape)

			if _, isHit := xs.Hit(); isHit {
				canvas.Write(x, y, color)
			}
		}
	}

	fmt.Println(canvas.ToPpm())

}
