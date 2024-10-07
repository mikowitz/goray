package main

import (
	"fmt"
	"math"

	g "github.com/mikowitz/goray/pkg"
)

func main() {
	floor := g.NewSphere()
	floor.SetTransform(g.Scaling(10, 0.01, 10))
	floor.Material.Color = g.NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0.0

	leftWall := g.NewSphere()
	leftWall.Transform = g.Translation(0, 0, 5).Mul(g.RotationY(-math.Pi / 4)).Mul(g.RotationX(math.Pi / 2)).Mul(g.Scaling(10, 0.01, 10))
	leftWall.SetMaterial(floor.Material)

	rightWall := g.NewSphere()
	rightWall.Transform = g.Translation(0, 0, 5).Mul(g.RotationY(math.Pi / 4)).Mul(g.RotationX(math.Pi / 2)).Mul(g.Scaling(10, 0.01, 10))
	rightWall.SetMaterial(floor.Material)

	middle := g.NewSphere()
	middle.SetTransform(g.Translation(-0.5, 1, 0.5))
	middle.Material.Color = g.NewColor(0.5, 0.1, 1)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := g.NewSphere()
	right.SetTransform(g.Translation(1.5, 0.5, -0.5).Mul(g.Scaling(0.5, 0.5, 0.5)))
	right.SetMaterial(middle.Material)

	left := g.NewSphere()
	left.SetTransform(g.Translation(-1.5, 0.33, -0.75).Mul(g.Scaling(0.33, 0.33, 0.33)))
	left.SetMaterial(middle.Material)

	world := g.NewWorld()
	world.LightSource = g.NewPointLight(g.NewPoint(-10, 10, -10), g.NewColor(1, 1, 1))
	world.Objects = []g.Sphere{
		floor, leftWall, rightWall,
		middle, right, left,
	}

	c := g.NewCamera(100, 16./9., math.Pi/3)
	c.Transform = g.NewViewTransform(
		g.NewPoint(0, 1.5, -5),
		g.NewPoint(0, 1, 0),
		g.NewVector(0, 1, 0),
	)

	canvas := c.Render(world)

	fmt.Println(canvas.ToPpm())
}
