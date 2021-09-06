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
