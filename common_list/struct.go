package common_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(vals []int) *ListNode {
	head := &ListNode{}
	p := head
	for k, v := range vals {
		p.Val = v
		if k < len(vals)-1 {
			p.Next = &ListNode{}
			p = p.Next
		}
	}
	return head
}

func (l *ListNode) Print() {
	for l != nil {
		fmt.Printf("%v ", l.Val)
		l = l.Next
	}
	fmt.Println()
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(vals *[]int) *TreeNode {
	if len(*vals) <= 0 {
		return nil
	}
	if (*vals)[0] == -1 {
		*vals = (*vals)[1:]
		return nil
	}
	res := &TreeNode{
		Val: (*vals)[0],
	}
	*vals = (*vals)[1:]
	res.Left = BuildTree(vals)
	res.Right = BuildTree(vals)
	return res
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

type PicNode struct {
	Val       int
	Neighbors []*PicNode
}
