package common_list

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

// 进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

// 思路：
// 快慢指针，反转链表
func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 快慢指针寻找中点
	q, s := head, head
	for q.Next != nil && q.Next.Next != nil {
		q = q.Next.Next
		s = s.Next
	}

	// 从中点反转链表
	mid := s.Next
	s.Next = nil

	p, n := mid, mid.Next
	mid.Next = nil
	for n != nil {
		p, n, n.Next = n, n.Next, p
	}

	// check两个链表是否相同
	for p != nil && head != nil {
		if p.Val != head.Val {
			return false
		}
		p = p.Next
		head = head.Next
	}

	return true
}

// 1 2 3 4 3 2 1
