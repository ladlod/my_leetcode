package week_contest

import "github.com/ladlod/leetcode/dp"

// 给你一个字符串 s ，请你找到 s 中两个 不相交回文子序列 ，使得它们长度的 乘积最大 。两个子序列在原字符串中如果没有任何相同下标的字符，则它们是 不相交 的。

// 请你返回两个回文子序列长度可以达到的 最大乘积 。

// 子序列 指的是从原字符串中删除若干个字符（可以一个也不删除）后，剩余字符不改变顺序而得到的结果。如果一个字符串从前往后读和从后往前读一模一样，那么这个字符串是一个 回文字符串 。

// 2 <= s.length <= 12
// s 只含有小写英文字母。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-product-of-the-length-of-two-palindromic-subsequences
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxProduct(s string) int {
	res := 0
	var dfs func(s, s1, s2 string)
	dfs = func(s, s1, s2 string) {
		if len(s) == 0 {
			if v := dp.LongestPalindromeSubseq(s1) * dp.LongestPalindromeSubseq(s2); v > res {
				res = v
			}
		} else {
			dfs(s[1:], s1+s[:1], s2)
			dfs(s[1:], s1, s2+s[:1])
		}
	}
	dfs(s, "", "")
	return res
}
