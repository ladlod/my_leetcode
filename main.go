package main

import (
	"github.com/ladlod/leetcode/common_list"
)

func main() {
	list := common_list.BuildList([]int{1, 2, 3, 4, 5})
	list.Print()
	list = common_list.ReverseList(list)
	list.Print()
}
