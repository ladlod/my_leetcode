package common_list

// 给你一个字符串 s，找到 s 中最长的回文子串。

// 思路：
// 暴力解法，以每个点为原点向两侧遍历，找到最长回文子串（分奇偶讨论），时间复杂度O(n^2)
func longestPalindrome(s string) string {
	res := innerLongestPalindrome(s, 0, 0)
	for i := 1; i < len(s); i++ {
		s1 := innerLongestPalindrome(s, i-1, i)
		s2 := innerLongestPalindrome(s, i, i)
		if len(s1) > len(s2) {
			if len(s1) > len(res) {
				res = s1
			}
		} else {
			if len(s2) > len(res) {
				res = s2
			}
		}
	}
	return res
}

func innerLongestPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
