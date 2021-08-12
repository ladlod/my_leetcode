package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	t := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	common_list.GameOfLife(t)
	fmt.Println(t)
}
