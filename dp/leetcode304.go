package dp

type NumMatrix struct {
	matrix [][]int
	square [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	res := NumMatrix{
		matrix: matrix,
		square: make([][]int, m),
	}

	for i := 0; i < m; i++ {
		res.square[i] = make([]int, n)
		for j := 0; j < n; j++ {
			x := 0
			if i > 0 {
				x = res.square[i-1][j]
			}
			y := 0
			if j > 0 {
				y = res.square[i][j-1]
			}
			z := 0
			if i > 0 && j > 0 {
				z = res.square[i-1][j-1]
			}
			res.square[i][j] = x + y - z + res.matrix[i][j]
		}
	}
	return res
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	x := 0
	if row1 > 0 {
		x = this.square[row1-1][col2]
	}
	y := 0
	if col1 > 0 {
		y = this.square[row2][col1-1]
	}
	z := 0
	if row1 > 0 && col1 > 0 {
		z = this.square[row1-1][col1-1]
	}
	return this.square[row2][col2] - x - y + z
}
