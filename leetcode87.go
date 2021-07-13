package main

import (
	"reflect"
)

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if !isEqual(s1, s2) {
		return false
	}
	if len(s1) != len(s2) {
		return false
	}
	l := len(s1)
	if l == 1 {
		return false
	}
	for i := 0; i < l-1; i++ {
		if isScramble(s1[0:i+1], s2[0:i+1]) && isScramble(s1[i+1:], s2[i+1:]) {
			return true
		}
		if isScramble(s1[0:i+1], s2[l-i-1:]) && isScramble(s1[i+1:], s2[0:l-i-1]) {
			return true
		}
	}

	return false
}

func isEqual(s1 string, s2 string) bool {
	map1, map2 := make(map[byte]int), make(map[byte]int)
	byte1 := []byte(s1)
	byte2 := []byte(s2)
	for _, b := range byte1 {
		_, ok := map1[b]
		if ok {
			map1[b]++
		} else {
			map1[b] = 1
		}
	}
	for _, b := range byte2 {
		_, ok := map2[b]
		if ok {
			map2[b]++
		} else {
			map2[b] = 1
		}
	}

	if reflect.DeepEqual(map1, map2) {
		return true
	}
	return false
}
