package interview

import (
	"fmt"
)

// 给你一个链表，从后向前 每 k 个节点一组进行翻转，请你返回翻转后的链表。

// k 是一个正整数，它的值小于或等于链表的长度。

// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

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
