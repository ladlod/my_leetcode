package dp

func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	if h <= 0 {
		return 0
	}
	res := make([][]int, h)
	res[0] = triangle[0]
	for i := 1; i < h; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				res[i][j] = res[i-1][j] + triangle[i][j]
			} else if j == i {
				res[i][j] = res[i-1][j-1] + triangle[i][j]
			} else {
				res[i][j] = min(res[i-1][j-1], res[i-1][j]) + triangle[i][j]
			}
		}
	}

	return min(res[h-1]...)
}
