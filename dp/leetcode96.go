package dp

func numTrees(n int) int {
	if n <= 0 {
		return 0
	}
	res := make([]int, n+1)
	res[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res[i] += res[i-1]
			} else {
				res[i] += res[j] * res[i-j-1]
			}
		}
	}

	return res[n]
}
