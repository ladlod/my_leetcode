package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	trees := generateTrees(n)

	for _, tree := range trees {
		tree.print()
		fmt.Println()
	}
}

func (t *TreeNode) print() {
	if t == nil {
		return
	}
	fmt.Printf("%d ", t.Val)
	if t.Left != nil {
		t.Left.print()
	} else if t.Right != nil {
		fmt.Print("null ")
	}
	if t.Right != nil {
		t.Right.print()
	} else if t.Left != nil {
		fmt.Print("null ")
	}
}

func generateTrees(n int) []*TreeNode {
	trees := make([]*TreeNode, 0)

	return trees
}
