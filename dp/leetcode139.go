package dp

import (
	"fmt"
	"strings"
)

// 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

// 说明：

// 拆分时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
// 示例 1：

// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-break
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 动态规划，dp[i]标记s到第i个字符是否可以由词典组成
// dp[0] = true
// 遍历词表
// dp[i] = dp[i-len(word)] && s[i-len(word):word] in == word
func WordBreak(s string, wordDict []string) bool {
	l := len(s)
	res := make([]bool, l+1)
	res[0] = true
	for i := 1; i <= l; i++ {
		for _, word := range wordDict {
			if strings.HasSuffix(s[:i], word) && (i-len(word) <= 0 || res[i-len(word)]) {
				res[i] = true
				break
			}
		}
	}
	fmt.Println(res)

	return res[l]
}
