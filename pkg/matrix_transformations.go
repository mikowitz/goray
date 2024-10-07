package goray

func Translation(x, y, z float64) Matrix {
	m := IdentityMatrix()
	m.Write(0, 3, x)
	m.Write(1, 3, y)
	m.Write(2, 3, z)
	return m
}
