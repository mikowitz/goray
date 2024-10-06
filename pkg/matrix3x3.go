package goray

type Matrix3x3 [9]float64

func NewMatrix3x3(m ...float64) Matrix3x3 {
	return Matrix3x3(m[:9])
}

func (m Matrix3x3) At(row, col int) float64 {
	return m[row*3+col]
}

func (m Matrix3x3) Submatrix(xrow, xcol int) Matrix2x2 {
	r := make([]float64, 0)
	for row := range 3 {
		for col := range 3 {
			if row != xrow && col != xcol {
				r = append(r, m.At(row, col))
			}
		}
	}
	return Matrix2x2(r)
}

func (m Matrix3x3) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

func (m Matrix3x3) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)
	if (row+col)%2 == 1 {
		return -minor
	}
	return minor
}

func (m Matrix3x3) Determinant() float64 {
	r := 0.0
	for col := range 3 {
		r += m.At(0, col) * m.Cofactor(0, col)
	}
	return r
}
