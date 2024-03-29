package dp

import (
	"strconv"
	"strings"
)

// 一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：

// "AAJF" ，将消息分组为 (1 1 10 6)
// "KJF" ，将消息分组为 (11 10 6)
// 注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。

// 给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。

// 题目数据保证答案肯定是一个 32 位 的整数。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/decode-ways
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/*
思路：
如果一个数字的首字符是0 或包含连续的两个0 则不存在解码方法
动态规划：从后向前
dp[n] == (
	if s[n] == '0' then 0
	num = num(s[n:n+2])
	if num == 10 || num == 20 then dp[n+2]
	else if num > 0 && num < 27 then dp[n+1] + dp[n+2]
	else dp[n+1]
)
*/

func numDecodings(s string) int {
	l := len(s)
	if l == 0 || s[:1] == "0" || strings.Contains(s, "00") {
		return 0
	}

	rtn := make([]int, l)
	if s[l-1:] == "0" {
		rtn[l-1] = 0
	} else {
		rtn[l-1] = 1
	}

	if l >= 2 {
		if s[l-2:l-1] == "0" {
			rtn[l-2] = 0
		} else {
			num, _ := strconv.ParseInt(s[l-2:], 10, 64)
			if num > 0 && num < 27 {
				rtn[l-2] = rtn[l-1] + 1
			} else {
				rtn[l-2] = rtn[l-1]
			}
		}
	}

	for i := l - 3; i >= 0; i-- {
		num, _ := strconv.ParseInt(s[i:i+2], 10, 64)
		if num%10 == 0 {
			if num > 0 && num < 27 && s[i+2] != '0' {
				rtn[i] = rtn[i+2]
			} else {
				return 0
			}
		} else {
			if num > 0 && num < 27 {
				rtn[i] = rtn[i+1] + rtn[i+2]
			} else {
				rtn[i] = rtn[i+1]
			}
		}
	}

	return rtn[0]
}
