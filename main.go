package main

import "github.com/ladlod/leetcode/common_list"

func main() {
	tree := common_list.BuildTreePost([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	tree.Print()
}
