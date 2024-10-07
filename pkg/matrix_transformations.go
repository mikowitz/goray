package goray

import "math"

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

func RotationX(radians float64) Matrix {
	m := IdentityMatrix()
	m.Write(1, 1, math.Cos(radians))
	m.Write(1, 2, -math.Sin(radians))
	m.Write(2, 1, math.Sin(radians))
	m.Write(2, 2, math.Cos(radians))
	return m
}

func RotationY(radians float64) Matrix {
	m := IdentityMatrix()
	m.Write(0, 0, math.Cos(radians))
	m.Write(0, 2, math.Sin(radians))
	m.Write(2, 0, -math.Sin(radians))
	m.Write(2, 2, math.Cos(radians))
	return m
}

func RotationZ(radians float64) Matrix {
	m := IdentityMatrix()
	m.Write(0, 0, math.Cos(radians))
	m.Write(0, 1, -math.Sin(radians))
	m.Write(1, 0, math.Sin(radians))
	m.Write(1, 1, math.Cos(radians))
	return m
}
