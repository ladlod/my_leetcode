package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	l := &ListNode{}
	p := l
	for i := 1; i <= 10; i++ {
		p.Val = i
		if i != 10 {
			p.Next = new(ListNode)
		}
		p = p.Next
	}
	l = reverseBetween(l, m, n)
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n || m < 1 {
		return head
	}
	p := &ListNode{Val: 0}
	rtn := p
	p.Next = head
	pnext, pre, start, end := &ListNode{}, &ListNode{}, &ListNode{}, &ListNode{}
	for i := 0; i <= n+1; i++ {
		if i == m-1 {
			pre = p    //0
			p = p.Next // 1
			pre.Next = nil
		} else if i == m {
			start = p           //1
			tmp := p.Next       //2
			pnext = p.Next.Next //3
			p.Next.Next = p     //2->1
			p = tmp             //2
			start.Next = nil
		} else if i > m && i < n {
			tmp := pnext.Next //4
			pnext.Next = p    //3->2
			p = pnext         //3
			pnext = tmp       //4
		} else if i == n {
			end = p   // 3
			p = pnext // 4
		} else if i == n+1 {
			pre.Next = end //0->3
			start.Next = p //1->4
		} else {
			p = p.Next
		}
	}
	return rtn.Next
}
