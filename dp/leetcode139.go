package dp

import (
	"fmt"
	"strings"
)

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
