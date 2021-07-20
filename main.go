package main

import "github.com/ladlod/leetcode/common_list"

func main() {
	tree := common_list.BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	tree.Print()
}
