package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var m, n int
	fmt.Scanf("%s", m)
	fmt.Scanf("%s", n)
	l := &ListNode{}
	p := l
	for i := 1; i <= 10; i++ {
		p.Val = i
		p = p.Next
	}
	reverseBetween(l, m, n)
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	p := head
	for i := 1; i < n; i++ {
		if i >= m-1 && i <= n {
			for p.Next != nil && p.Next.Next != nil {
				tmp1 := p.Next
				tmp2 := p.Next.Next
				p.Next = tmp2
				tmp1.Next = tmp2.Next
				tmp2.Next = tmp1
				p = tmp1
				i++
			}
		}
		p = p.Next
	}
	return head
}
