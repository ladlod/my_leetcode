package interview

import (
	"math"
	"sort"
)

func main() {
	// fmt.Println(parentheses("]"))
	// fmt.Println(anagram("anagram", "nagaramb"))
	// fmt.Println(findMax([]int{2, 5, 5}, 1))
}

// Validate Parentheses (string contains chars '(', ')', '[', ']', '{', '}')	s = "()[]{}"
func parentheses(str string) bool {
	var stack []rune
	for _, c := range str {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// Valid Anagram of 2 strings	s = "anagram", t = "nagaram"
func anagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	check := make(map[rune]int)

	for _, c := range s {
		check[c]++
	}

	for _, c := range t {
		check[c]--
	}

	for _, v := range check {
		if v != 0 {
			return false
		}
	}
	return true
}

/*
Given n integers {a1, a2, ... an} and a target sum k. Find an integer sequence {b1, b2, ... bn} such that
	1)	bi <= ai
	2)	b1 + b2 + ... + bn = k
	3)	Minimize max(bi)

Example:
	Input: a[] = [2,1,5,6,2,3], k = 13
	Solution: b[] = [2,1,3,3,2,2] max(bi) = 3
	Output: 3

solution:
	1. if sum(a) < k return 0
	2. sum(a) - k
	[1,2,2,3,5,6]
*/
func findMax(slice []int, k int) int {
	sum := 0
	min := math.MaxInt32
	for _, v := range slice {
		sum += v
		if v < min {
			min = v
		}
	}
	if sum < k {
		return 0
	} else if sum == k {
		return min
	} else {
		sub := sum - k
		sort.Ints(slice)
		for i := len(slice) - 1; i > 0; i-- {
			for slice[i] > slice[i-1] && sub > 0 {
				for j := i; j < len(slice) && sub > 0; j++ {
					slice[j]--
					sub--
				}
			}
		}
		return slice[len(slice)-1] - sub/len(slice)
	}
}
