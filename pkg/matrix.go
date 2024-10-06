package goray

type Matrix [16]float64
type Matrix3x3 [9]float64
type Matrix2x2 [4]float64

func NewMatrix(m ...float64) Matrix {
	return Matrix(m[:16])
}

func NewMatrix3x3(m ...float64) Matrix3x3 {
	return Matrix3x3(m[:9])
}

func NewMatrix2x2(m ...float64) Matrix2x2 {
	return Matrix2x2(m[:4])
}

func (m Matrix) At(row, col int) float64 {
	return m[row*4+col]
}

func (m Matrix3x3) At(row, col int) float64 {
	return m[row*3+col]
}

func (m Matrix2x2) At(row, col int) float64 {
	return m[row*2+col]
}
