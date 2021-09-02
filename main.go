package main

import (
	"github.com/ladlod/leetcode/common_list"
)

func main() {
	list1 := common_list.BuildList([]int{1})
	list2 := common_list.BuildList([]int{0})
	list := common_list.MergeKLists([]*common_list.ListNode{list1, list2})
	list.Print()
}
