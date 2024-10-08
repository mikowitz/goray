package main

import (
	"fmt"
	"math"

	g "github.com/mikowitz/goray/pkg"
)

func main() {
	floorPattern1 := g.NewStripePattern(g.NewColor(1, 0.9, 0.9), g.NewColor(0.0, 0.5, 0.1))
	floorPattern2 := g.NewStripePattern(g.NewColor(1, 0.9, 0.9), g.NewColor(0.0, 0.5, 0.1))
	floorPattern2.SetTransform(g.RotationY(math.Pi / 2))
	floorPattern := g.BlendedPattern{A: &floorPattern1, B: &floorPattern2}
	floorPattern.SetTransform(g.Scaling(0.25, 0.25, 0.25))
	floor := g.NewPlane()
	floor.Material.Pattern = &floorPattern
	floor.Material.Specular = 0.0

	wall := g.NewPlane()
	wall.SetTransform(g.Translation(0, 0, 4).Mul(g.RotationY(math.Pi / 4)).Mul(g.RotationX(math.Pi / 2)))
	wall.Material.Pattern = &floorPattern1

	otherWall := g.NewPlane()
	otherWall.SetTransform(g.Translation(0, 0, 6).Mul(g.RotationY(-math.Pi / 4)).Mul(g.RotationX(math.Pi / 2)))
	// otherWall.SetMaterial(floor.Material)

	middle := g.NewSphere()
	middle.SetTransform(g.Translation(-0.5, 1, 0.5))
	purple := g.NewCheckersPattern(g.NewColor(0.5, 0.1, 1), g.NewColor(0, 0.5, 1))
	purple.SetTransform(g.Scaling(0.25, 0.25, 0.25))
	middle.Material.Pattern = &purple
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
	world.Objects = []g.Shape{}

	c := g.NewCamera(200, 16./9., math.Pi/3)
	c.Transform = g.NewViewTransform(
		g.NewPoint(0, 1.5, -5),
		g.NewPoint(0, 1, 0),
		g.NewVector(0, 1, 0),
	)

	canvas := c.Render(world)

	fmt.Println(canvas.ToPpm())
}
