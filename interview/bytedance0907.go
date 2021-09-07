package interview

import (
	"fmt"
)

type listNode struct {
	Val  int
	Next *listNode
}

// 初始化链表
func initList(vals []int) *listNode {
	head := new(listNode)
	p := head
	for _, v := range vals {
		p.Next = &listNode{
			Val: v,
		}
		p = p.Next
	}
	return head.Next
}

// 打印链表
func (l *listNode) print() {
	for ; l != nil; l = l.Next {
		fmt.Printf("%d ", l.Val)
	}
	fmt.Println()
}

func reverseKList(head *listNode, k int) *listNode {
	pre := new(listNode)
	pre.Next = head

	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	if l < k {
		return head
	}

	j := l % k
	q := pre
	for j > 0 {
		q = q.Next
		j--
	}

	// q = 1

	s := q
	for i := k; i > 0; i-- {
		q = q.Next
	}

	// s = 1, q = 4

	for q != nil {
		next := q.Next
		q.Next = nil
		s.Next.Next, s.Next, q = next, reverseList(s.Next), s.Next

		for i := k; i > 0 && s != nil && q != nil; i-- {
			s = s.Next
			q = q.Next
		}
	}
	return pre.Next
}

// 反转链表
func reverseList(head *listNode) *listNode {
	if head.Next == nil {
		return head
	}

	p := head
	q := p.Next
	p.Next = nil
	for q != nil {
		q.Next, p, q = p, q, q.Next
	}
	return p
}
