package common_list

import (
	"fmt"
	"reflect"
)

func FindSubstring(s string, words []string) []int {
	res := make([]int, 0)
	if len(words) == 0 {
		return res
	}
	l := len(words[0])

	juegeMap := make(map[string]int)
	for _, word := range words {
		juegeMap[word]++
	}

	for i := 0; i <= len(s)-l*len(words); i++ {
		if judge30(s[i:i+l*len(words)], juegeMap, l) {
			res = append(res, i)
		}
	}

	return res
}

func judge30(s string, juegeMap map[string]int, l int) bool {
	tmpMap := make(map[string]int)
	for i := 0; i < len(s); i += l {
		tmpMap[s[i:i+l]]++
	}
	fmt.Println(tmpMap)
	return reflect.DeepEqual(tmpMap, juegeMap)
}
