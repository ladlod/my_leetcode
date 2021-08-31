package common_list

import "fmt"

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 示例 4:

// 输入: s = ""
// 输出: 0

// 思路：滑动窗口 维护一个map存储当前所有的字符串及对应的index，顺序遍历，每次发现冲突时减掉冲突及之前的字符串

func LengthOfLongestSubstring(s string) int {
	sMap := make(map[byte]int)
	start, end, res := 0, 0, 0
	for i := range s {
		if index, ok := sMap[s[i]]; ok {
			for j := start; j <= index; j++ {
				delete(sMap, s[j])
			}
			start = index + 1
		}
		end = i
		sMap[s[i]] = i
		if end-start+1 > res {
			res = end - start + 1
		}

		fmt.Println(s[start : end+1])
	}

	return res
}
