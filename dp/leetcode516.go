package dp

// 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。

// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindromic-subsequence
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路:
// 动态规划
// dp[i][j]记录从i到j的最长回文字符串
// dp[i][j] = if i == j then 1
//            elif s[i] == s[j] then dp[i+1][j-1]+2
//			  else max(dp[i-1][j], dp[i][j-1])
func LongestPalindromeSubseq(s string) int {
	l := len(s)
	if l <= 0 {
		return 0
	}
	res := make([][]int, l)
	for i := l - 1; i >= 0; i-- {
		res[i] = make([]int, l)
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
					res[i][j] = max(res[i][j-1], res[i+1][j])
				}
			}
		}
	}

	return res[0][l-1]
}
