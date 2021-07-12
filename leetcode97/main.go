package main

import (
	"fmt"
)

func main() {
	var a, b, c string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	fmt.Scanf("%s", &c)
	fmt.Println(isInterleave(a, b, c))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	} else {
		res := make([][]bool, len(s1)+1)
		for i := 0; i <= len(s1); i++ {
			res[i] = make([]bool, len(s2)+1)
			res[i][0] = s1[:i] == s3[:i]
		}
		for j := 0; j <= len(s2); j++ {
			res[0][j] = s2[:j] == s3[:j]
		}

		for i := 0; i < len(s1); i++ {
			for j := 0; j < len(s2); j++ {
				if res[i][j+1] {
					if s1[i] == s3[i+j+1] {
						res[i+1][j+1] = true
					}
				} else if res[i+1][j] {
					if s2[j] == s3[i+j+1] {
						res[i+1][j+1] = true
					}
				} else {
					res[i+1][j+1] = false
				}
			}
		}
		return res[len(s1)][len(s2)]
	}
}
