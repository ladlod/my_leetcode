package daily

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(vals []int) *ListNode {
	if len(vals) <= 0 {
		return nil
	}
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
