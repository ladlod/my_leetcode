package dp

func LongestPalindromeSubseq(s string) int {
	l := len(s)
	if l <= 0 {
		return 0
	}
	res := make([][]int, l)
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			if i == j {
				res[i][j] = 1
			} else if i == j-1 {
				if s[i] == s[j] {
					res[i][j] = 2
				} else {
					res[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					res[i][j] = res[i+1][j-1] + 2
				} else {
					res[i][j] = res[i+1][j-1]
				}
			}
		}
	}

	return res[0][l-1]
}
