package goray

func Translation(x, y, z float64) Matrix {
	m := IdentityMatrix()
	m.Write(0, 3, x)
	m.Write(1, 3, y)
	m.Write(2, 3, z)
	return m
}

func Scaling(x, y, z float64) Matrix {
	m := IdentityMatrix()
	m.Write(0, 0, x)
	m.Write(1, 1, y)
	m.Write(2, 2, z)
	return m
}
