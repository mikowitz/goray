package goray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DemoShape struct {
	Material  Material
	Transform Matrix
	SavedRay  Ray
}

func NewDemoShape() DemoShape {
	return DemoShape{
		Material:  NewMaterial(),
		Transform: IdentityMatrix(),
	}
}

func (ds *DemoShape) GetMaterial() Material {
	return ds.Material
}

func (ds *DemoShape) SetMaterial(m Material) {
	ds.Material = m
}

func (ds *DemoShape) GetTransform() Matrix {
	return ds.Transform
}

func (ds *DemoShape) SetTransform(m Matrix) {
	ds.Transform = m
}

func (ds *DemoShape) GetSavedRay() Ray {
	return ds.SavedRay
}

func (ds *DemoShape) LocalIntersect(r Ray) Intersections {
	ds.SavedRay = r
	return Intersections{}
}

func (ds *DemoShape) LocalNormalAt(p Point) Vector {
	return NewVector(p.x, p.y, p.z)
}

func TestShapeTransform(t *testing.T) {
	t.Run("the default transformation", func(t *testing.T) {
		s := NewDemoShape()
		assert.True(t, MatricesEqual(s.Transform, IdentityMatrix()))
	})

	t.Run("assigning a transformation", func(t *testing.T) {
		s := NewDemoShape()
		s.SetTransform(Translation(2, 3, 4))
		assert.True(t, MatricesEqual(s.Transform, Translation(2, 3, 4)))
	})
}

func TestShapeMaterial(t *testing.T) {
	t.Run("the default material", func(t *testing.T) {
		s := NewDemoShape()
		assert.Equal(t, s.Material, NewMaterial())
	})

	t.Run("assigning a material", func(t *testing.T) {
		s := NewDemoShape()
		m := NewMaterial()
		m.Ambient = 1.0
		s.SetMaterial(m)
		assert.Equal(t, s.Material, m)
	})
}

func TestIntersection(t *testing.T) {
	t.Run("of a scaled shape", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewDemoShape()
		s.SetTransform(Scaling(2, 2, 2))
		_ = r.Intersect(&s)

		assert.True(t, TuplesEqual(s.GetSavedRay().Origin, NewPoint(0, 0, -2.5)))
		assert.True(t, TuplesEqual(s.GetSavedRay().Direction, NewVector(0, 0, 0.5)))
	})

	t.Run("of a translated shape", func(t *testing.T) {
		r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
		s := NewDemoShape()
		s.SetTransform(Translation(5, 0, 0))
		_ = r.Intersect(&s)

		assert.True(t, TuplesEqual(s.GetSavedRay().Origin, NewPoint(-5, 0, -5)))
		assert.True(t, TuplesEqual(s.GetSavedRay().Direction, NewVector(0, 0, 1)))
	})
}

func TestNormal(t *testing.T) {
	t.Run("on a translated shape", func(t *testing.T) {
		s := NewDemoShape()
		s.SetTransform(Translation(0, 1, 0))
		n := NormalAt(&s, NewPoint(0, 1.70711, -0.70711))

		assert.True(t, TuplesEqual(n, NewVector(0, 0.70711, -0.70711)))
	})
}
