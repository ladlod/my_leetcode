package interview

import "strconv"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	sMap := make(map[string][]string)
	for _, str := range strs {
		key := hashKey(str)
		if _, ok := sMap[key]; !ok {
			sMap[key] = make([]string, 0)
		}
		sMap[key] = append(sMap[key], str)
	}

	for _, r := range sMap {
		res = append(res, r)
	}
	return res
}

func hashKey(str string) string {
	res := make([]byte, 0)
	sMap := make(map[byte]int)
	for _, v := range []byte(str) {
		sMap[v]++
	}
	for _, v := range []byte(alphabet) {
		if sMap[v] > 0 {
			res = append(res, v)
			res = append(res, []byte(strconv.Itoa(sMap[v]))...)
		}
	}
	return string(res)
}
