package common_list

import "fmt"

// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/regular-expression-matching
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 动态规划
// dp[i][j] 表示s[i]与p[j]是否匹配
// if s[i] == p[j] || s[i] == '.' || p[j] == '.': dp[i][j] = dp[i-1][j-1]
func IsMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
		for j := 0; j <= len(p); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				if p[j-1] == '*' && j > 1 {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = false
				}
			} else if j == 0 {
				if s[i-1] == '*' && i > 1 {
					dp[i][j] = dp[i-2][j]
				} else {
					dp[i][j] = false
				}
			} else {
				if s[i-1] == p[j-1] || s[i-1] == '.' || p[j-1] == '.' {
					dp[i][j] = dp[i-1][j-1]
				} else if s[i-1] == '*' && i > 1 {
					dp[i][j] = dp[i-2][j] || ((p[j-1] == s[i-2] || p[j-1] == '.' || s[i-2] == '.') && dp[i][j-1])
				} else if p[j-1] == '*' && j > 1 {
					dp[i][j] = dp[i][j-2] || ((p[j-2] == s[i-1] || p[j-2] == '.' || s[i-1] == '.') && dp[i-1][j])
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	fmt.Println(dp)

	return dp[len(s)][len(p)]
}
