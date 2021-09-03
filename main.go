package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	tree := &common_list.TreeNode{
		Val: 1,
		Left: &common_list.TreeNode{
			Val: 2,
			Left: &common_list.TreeNode{
				Val: 4,
			},
		},
		Right: &common_list.TreeNode{
			Val: 3,
		},
	}
	codeC := common_list.Constructor()
	str := codeC.Serialize(tree)
	fmt.Println(str)
	newTree := codeC.Deserialize(str)
	newTree.Print()
}
