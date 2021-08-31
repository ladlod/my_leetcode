package main

import (
	"github.com/ladlod/leetcode/common_list"
)

func main() {
	list := common_list.BuildList([]int{1})
	list.Print()
	common_list.ReorderList(list)
	// fmt.Println(common_list.LengthOfLongestSubstring("abcabcbb"))
}
