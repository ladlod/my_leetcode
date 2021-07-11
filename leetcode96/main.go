package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(numTrees(n))
}

func numTrees(n int) int {
	res := make([]int, n)
	for i := 0; i <= n; i++ {
		res[i] = res[i-1] * res[n-i]
	}

	return res[n-1]
}
