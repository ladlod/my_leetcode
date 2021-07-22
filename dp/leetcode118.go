package dp

func generateYanghui(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 {
				res[i][j] = res[i-1][j]
			} else if j == i {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}
