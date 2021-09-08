package main

import (
	"github.com/ladlod/leetcode/common_list"
)

func main() {
	list := common_list.BuildList([]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 5})
	common_list.DeleteDuplicates(list).Print()
}
