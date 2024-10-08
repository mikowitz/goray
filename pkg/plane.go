package goray

import "math"

type Plane struct {
	Material  Material
	Transform Matrix
	SavedRay  Ray
}

func NewPlane() Plane {
	return Plane{
		Material:  NewMaterial(),
		Transform: IdentityMatrix(),
	}
}

func (p *Plane) GetMaterial() Material {
	return p.Material
}

func (p *Plane) SetMaterial(m Material) {
	p.Material = m
}

func (p *Plane) GetTransform() Matrix {
	return p.Transform
}

func (p *Plane) SetTransform(m Matrix) {
	p.Transform = m
}

func (p *Plane) GetSavedRay() Ray {
	return p.SavedRay
}

func (p *Plane) LocalIntersect(r Ray) Intersections {
	if math.Abs(r.Direction.y) < 0.00001 {
		return Intersections{}
	}
	// return Intersections{}
	t := -r.Origin.y / r.Direction.y
	return Intersections{
		NewIntersection(t, p),
	}
}

func (p *Plane) LocalNormalAt(_ Point) Vector {
	return NewVector(0, 1, 0)
}
