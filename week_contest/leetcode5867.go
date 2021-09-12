package week_contest

// 给你一个下标从 0 开始的字符串 word 和一个字符 ch 。找出 ch 第一次出现的下标 i ，反转 word 中从下标 0 开始、直到下标 i 结束（含下标 i ）的那段字符。如果 word 中不存在字符 ch ，则无需进行任何操作。

// 例如，如果 word = "abcdefd" 且 ch = "d" ，那么你应该 反转 从下标 0 开始、直到下标 3 结束（含下标 3 ）。结果字符串将会是 "dcbaefd" 。
// 返回 结果字符串 。

func reversePrefix(word string, ch byte) string {
	bWord := []byte(word)
	for i := range bWord {
		if bWord[i] == ch {
			for j := 0; j < i; {
				bWord[i], bWord[j] = bWord[j], bWord[i]
				i--
				j++
			}
			break
		}
	}
	return string(bWord)
}
