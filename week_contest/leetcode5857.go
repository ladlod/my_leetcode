// 给你一个二进制字符串 binary 。 binary 的一个 子序列 如果是 非空 的且没有 前导 0 （除非数字是 "0" 本身），那么它就是一个 好 的子序列。

// 请你找到 binary 不同好子序列 的数目。

// 比方说，如果 binary = "001" ，那么所有 好 子序列为 ["0", "0", "1"] ，所以 不同 的好子序列为 "0" 和 "1" 。 注意，子序列 "00" ，"01" 和 "001" 不是好的，因为它们有前导 0 。
// 请你返回 binary 中 不同好子序列 的数目。由于答案可能很大，请将它对 10^9 + 7 取余 后返回。

// 一个 子序列 指的是从原数组中删除若干个（可以一个也不删除）元素后，不改变剩余元素顺序得到的序列。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-unique-good-subsequences
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 动态规划
// dp[i][0] = dp[i-1][0] + (binary[i] == '0' * dp[i-1][1] + 1) 计算所有以0开始不重复的子序列
// dp[i][1] = dp[i-1][1] + (binary[i] == '1' * dp[i-1][0] + 1) 计算所有以1开始不重复的子序列

package week_contest

import (
	"math"
)

func NumberOfUniqueGoodSubsequences(binary string) int {
	dp := make([][]int, len(binary))
	for i := range binary {
		dp[i] = make([]int, 2)
	}

	var has bool = false
	for i, j := len(binary)-1, 0; i >= 0; i, j = i-1, j+1 {
		if j == 0 {
			if binary[i] == '0' {
				dp[j][0]++
				has = true
			} else {
				dp[j][1]++
			}
		} else {
			if binary[i] == '0' {
				dp[j][0] = (dp[j-1][0] + dp[j-1][1] + 1) % (int(math.Pow(10, 9)) + 7)
				dp[j][1] = dp[j-1][1]
				has = true
			} else {
				dp[j][0] = dp[j-1][0]
				dp[j][1] = (dp[j-1][1] + dp[j-1][0] + 1) % (int(math.Pow(10, 9)) + 7)
			}
		}
		// fmt.Println(dp)
	}

	if has {
		return dp[len(binary)-1][1] + 1
	} else {
		return dp[len(binary)-1][1]
	}
}
