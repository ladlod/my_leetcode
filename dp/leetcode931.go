package dp

func minFallingPathSum(matrix [][]int) int {
	l := len(matrix)
	if l <= 0 {
		return 0
	}
	res := make([][]int, l)
	res[0] = matrix[0]
	for i := 1; i < l; i++ {
		res[i] = make([]int, l)
		for j := 0; j < l; j++ {
			if j == 0 {
				res[i][j] = min(res[i-1][j], res[i-1][j+1]) + matrix[i][j]
			} else if j == l-1 {
				res[i][j] = min(res[i-1][j-1], res[i-1][j]) + matrix[i][j]
			}
			res[i][j] = min(res[i-1][j-1], res[i-1][j], res[i-1][j+1]) + matrix[i][j]
		}
	}
	return min(res[l-1]...)
}
