package lintcode

import "bytes"

// 描述
// 给出两个字符串, 你需要修改第一个字符串，将所有与第二个字符串中相同的字符删除, 并且第二个字符串中不同的字符与第一个字符串的不同字符连接

/*
思路：
求两个字符串的异或
*/

func concatenetedString(s1 string, s2 string) string {
	m1 := make(map[rune]int64)
	m2 := make(map[rune]int64)

	for _, c := range s1 {
		m1[c]++
	}

	for _, c := range s2 {
		m2[c]++
	}

	buf := new(bytes.Buffer)
	for _, c := range s1 {
		if m2[c] == 0 {
			buf.WriteRune(c)
		}
	}

	for _, c := range s2 {
		if m1[c] == 0 {
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
