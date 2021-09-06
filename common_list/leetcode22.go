package common_list

import "fmt"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 有效括号组合需满足：左括号必须以正确的顺序闭合。

// 思路：
// 回溯法：
// 括号+'('时，i++, 括号+')' i--,n--
// i > 0时，括号可以+'('和')'，i=0时，括号仅可以+'('
func GenerateParenthesis(n int) []string {
	return innnerGenerateParenthesis(0, n)
}

func innnerGenerateParenthesis(i, n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
		return res
	}
	if i < n {
		vs := innnerGenerateParenthesis(i+1, n)
		for _, v := range vs {
			res = append(res, "("+v)
		}
	}
	if i > 0 {
		vs := innnerGenerateParenthesis(i-1, n-1)
		for _, v := range vs {
			res = append(res, ")"+v)
		}
	}
	fmt.Println(i, n, res)
	return res
}

// 思路：
// 动态规划：
// 记录dp[i]为n=i时的所有括号组合
// 那么，对于dp[1]，其可能的组合为 '()'
//      对于dp[2]，其可能的组合为 '(())','()()'
//      对于dp[3]，其可能的组合为 ‘()(())', '()()()', '(())()', '(()())', '((()))'
// 所以，可以推导出递推公式，dp[i] = '('+dp[p]+')'+dp[q]， 其中p+q = i-1

func GenerateParenthesisV2(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = []string{""}

	for i := 1; i <= n; i++ {
		dp[i] = make([]string, 0)
		for p := 0; p <= i-1; p++ {
			for _, v1 := range dp[p] {
				for _, v2 := range dp[i-1-p] {
					dp[i] = append(dp[i], "("+v1+")"+v2)
				}
			}
		}
	}
	return dp[n]
}
