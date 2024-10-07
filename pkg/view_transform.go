package goray

func NewViewTransform(from, to Point, up Vector) Matrix {
	forward := to.Sub(from).Normalize()
	left := forward.Cross(up.Normalize())
	trueUp := left.Cross(forward)

	m := NewMatrix(
		left.x, left.y, left.z, 0,
		trueUp.x, trueUp.y, trueUp.z, 0,
		-forward.x, -forward.y, -forward.z, 0,
		0, 0, 0, 1,
	)

	return m.Mul(Translation(-from.x, -from.y, -from.z))
}
