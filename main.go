package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	tree := common_list.BuildTree(&[]int{1, 2, 3, -1, -1, 4, -1, -1, 5, -1, 6, -1, -1})
	tree.Print()
	fmt.Println()
	common_list.Flatten(tree)
	tree.Print()
}
