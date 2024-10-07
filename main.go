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
	shape := g.NewSphere()
	m := g.NewMaterial()
	m.Color = g.NewColor(1, 0.2, 1)
	shape.SetMaterial(m)

	light := g.NewPointLight(g.NewPoint(-10, 10, -10), g.NewColor(1, 1, 1))

	for y := range canvas.Height {
		worldY := half - pixelSize*float64(y)

		for x := range canvas.Width {
			worldX := -half + pixelSize*float64(x)

			position := g.NewPoint(worldX, worldY, wallZ)

			r := g.NewRay(rayOrigin, position.Sub(rayOrigin).Normalize())
			xs := r.Intersect(shape)

			if hit, isHit := xs.Hit(); isHit {
				point := r.At(hit.T)
				normal := hit.Object.NormalAt(point)
				eye := r.Direction.Neg()

				color := hit.Object.Material.Lighting(light, point, eye, normal)
				canvas.Write(x, y, color)
			}
		}
	}

	fmt.Println(canvas.ToPpm())

}
