package common_list

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

func swapPairs(head *ListNode) *ListNode {
	pre := new(ListNode)
	pre.Next = head
	p := pre

	for p != nil && p.Next != nil && p.Next.Next != nil {
		p.Next, p.Next.Next, p.Next.Next.Next = p.Next.Next, p.Next.Next.Next, p.Next
		p = p.Next.Next
	}
	return pre.Next
}
