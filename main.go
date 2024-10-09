package main

import (
	"fmt"
	"math"

	g "github.com/mikowitz/goray/pkg"
)

func main() {
	// floorPattern := g.NewSolidPattern(g.NewColor(0.9, 0.9, 0.9))
	floorPattern := g.NewCheckersPattern(g.NewColor(0, 0, 0), g.NewColor(1, 1, 1))
	floor := g.NewPlane()
	fm := g.NewMaterial()
	fm.Pattern = &floorPattern
	fm.Reflective = 0.25
	floor.SetMaterial(fm)

	wallPattern := g.NewStripePattern(g.NewColor(0.9, 0.9, 0.9), g.NewColor(0.2, 0.2, 0.2))
	wallPattern.SetTransform(g.Scaling(0.33, 0.33, 0.33))

	wall := g.NewPlane()
	wall.SetTransform(g.Translation(0, 0, 10).Mul(g.RotationY(math.Pi / 3)).Mul(g.RotationZ(math.Pi / 2)).Mul(g.RotationX(math.Pi / 2)))
	wall.Material.Pattern = &wallPattern

	otherWall := g.NewPlane()
	otherWall.SetTransform(g.Translation(0, 0, 10).Mul(g.RotationY(-math.Pi / 3)).Mul(g.RotationZ(math.Pi / 2)).Mul(g.RotationX(math.Pi / 2)))
	otherWall.Material.Pattern = &wallPattern

	red := g.NewSolidPattern(g.NewColor(1, 0, 0.5))
	sphere := g.NewSphere()
	sphere.SetTransform(g.Translation(0, 1, 0))
	m := g.NewMaterial()
	m.Pattern = &red
	m.Shininess = 50.0
	sphere.SetMaterial(m)

	world := g.NewWorld()
	world.LightSource = g.NewPointLight(g.NewPoint(-10, 10, -10), g.NewColor(1, 1, 1))
	world.Objects = []g.Shape{&floor, &wall, &otherWall, &sphere}

	c := g.NewCamera(200, 16./9., math.Pi/3)
	c.Transform = g.NewViewTransform(
		g.NewPoint(0, 1.5, -5),
		g.NewPoint(0, 1, 0),
		g.NewVector(0, 1, 0),
	)

	canvas := c.Render(world)

	fmt.Println(canvas.ToPpm())
}
