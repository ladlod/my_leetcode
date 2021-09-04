package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	fmt.Println(common_list.FindDiagonalOrderV2([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}))
}
