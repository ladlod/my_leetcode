package common_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Print() {
	if t == nil {
		return
	}
	fmt.Printf("%d ", t.Val)
	if t.Left != nil {
		t.Left.Print()
	} else if t.Right != nil {
		fmt.Print("null ")
	}
	if t.Right != nil {
		t.Right.Print()
	} else if t.Left != nil {
		fmt.Print("null ")
	}
}
