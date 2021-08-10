package lintcode

// 给出一个字符串。找到字符串中第一个不重复的字符然后返回它的下标。如果不存在这样的字符，返回 -1
func FirstUniqChar(s string) int {
	cMap := make(map[rune]int)
	for _, c := range s {
		cMap[c]++
	}
	for k, c := range s {
		if cMap[c] == 1 {
			return k
		}
	}
	return -1
}
