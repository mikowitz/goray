package goray

type Shape interface {
	GetMaterial() Material
	SetMaterial(m Material)
	GetTransform() Matrix
	SetTransform(m Matrix)
	GetSavedRay() Ray
	LocalIntersect(r Ray) Intersections
	LocalNormalAt(p Point) Vector
}

func NormalAt(shape Shape, point Point) Vector {
	objectPoint := shape.GetTransform().Inverse().Mult(point)
	objectNormal := shape.LocalNormalAt(objectPoint)
	worldNormal := shape.GetTransform().Inverse().Transpose().Mult(objectNormal)
	worldNormal.w = 0.0
	return worldNormal.Normalize()
}
