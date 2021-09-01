package common_list

import "fmt"

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

//

// 注意：

// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。

// 示例 1：

// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-window-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 思路1：滑动窗口，先找到第一个包含子串t的窗口，之后向右滑动，滑动过程中若判定再次满足包含子串t，从左侧缩减子串长度至满足包含子串t的最小窗口，继续滑动并循环以上步骤
func MinWindowV1(s string, t string) string {
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)

	for i := range t {
		tMap[t[i]]++
	}

	l, r := 0, 0
	i := 0
	for ; !sHasT(sMap, tMap) && i < len(s); i++ {
		sMap[s[i]]++
		r++
	}
	if !sHasT(sMap, tMap) {
		return ""
	}
	res := string(s[l:r])

	for ; i <= len(s); i++ {
		for {
			sMap[s[l]]--
			if sHasT(sMap, tMap) {
				l++
				res = string(s[l:r])
			} else {
				sMap[s[l]]++
				break
			}
		}

		if i != len(s) {
			sMap[s[l]]--
			l++
			sMap[s[r]]++
			r++
		}
		fmt.Println(s[l:r])
	}

	return res
}

func sHasT(smap, tmap map[byte]int) bool {
	for k, v := range tmap {
		if smap[k] < v {
			return false
		}
	}
	return true
}
