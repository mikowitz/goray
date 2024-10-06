package goray

type Matrix2x2 [4]float64

func NewMatrix2x2(m ...float64) Matrix2x2 {
	return Matrix2x2(m[:4])
}

func (m Matrix2x2) At(row, col int) float64 {
	return m[row*2+col]
}

func (m Matrix2x2) Determinant() float64 {
	return m[0]*m[3] - m[1]*m[2]
}
