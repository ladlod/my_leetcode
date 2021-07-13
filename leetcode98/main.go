package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	} else if (root.Left != nil && root.Left.Val >= root.Val) ||
		(root.Right != nil && root.Right.Val <= root.Val) {
		fmt.Println(root.Left.Val, root.Right.Val)
		return false
	}
	return isValidBST(root.Right) && isValidBST(root.Left)
}
