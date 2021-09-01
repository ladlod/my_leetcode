package main

import (
	"github.com/ladlod/leetcode/common_list"
)

func main() {
	list := common_list.BuildList([]int{1, 2, 3, 4, 5, 6})
	list = common_list.OddEvenList(list)
	list.Print()
}
