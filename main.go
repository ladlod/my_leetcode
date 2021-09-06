package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	fmt.Println(common_list.HasCicle(map[int][]int{
		1: {2, 3},
		2: {3, 4},
	}))
}
