package common_list

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

// 思路：
// 用一个栈（栈长）x记录括号数，每当遇到(，x++，每当遇到)x--, x < 0时，记录已遍历的最大长度，
// 这样计算会略过x始终大于0的情况，所以以同样的方式再反向遍历一次即可
func LongestValidParentheses(s string) int {
	l := len(s)
	res, stack, tmp := 0, 0, 0
	for i := 0; i < l; i++ {
		tmp++
		if s[i] == '(' {
			stack++
		} else if s[i] == ')' {
			stack--
		}
		if stack < 0 {
			tmp = 0
			stack = 0
		} else if stack == 0 {
			if tmp > res {
				res = tmp
			}
		}
	}

	stack, tmp = 0, 0
	for i := l - 1; i >= 0; i-- {
		tmp++
		if s[i] == ')' {
			stack++
		} else if s[i] == '(' {
			stack--
		}
		if stack < 0 {
			tmp = 0
			stack = 0
		} else if stack == 0 {
			if tmp > res {
				res = tmp
			}
		}
	}

	return res
}
