package common_list

// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

// k 是一个正整数，它的值小于或等于链表的长度。

// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 双指针，快指针与慢指针间隔k，每次反转快慢指针之间的链表
func ReverseKGroup(head *ListNode, k int) *ListNode {
	pre := new(ListNode)
	pre.Next = head

	// step1: gen ptr
	q, s := pre, pre
	for i := 1; i <= k; i++ {
		if q.Next == nil {
			return head
		}
		q = q.Next
	}

	// step2: reverse
	for q != nil {
		start, end, next := s.Next, q, q.Next // 记录链表开始，结束，及下一个值
		end.Next = nil                        // 截断链表
		s.Next = ReverseList(start)           // 反转链表
		start.Next = next                     // 整合

		s, q = start, start
		for i := 1; i <= k; i++ {
			if q == nil {
				return pre.Next
			}
			q = q.Next
		}
		pre.Print()
	}
	return pre.Next
}

// 反转链表 递归
func ReverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	end := ReverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return end
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p, n := head, head.Next
	head.Next = nil
	for n != nil {
		n.Next, p, n = p, n, n.Next
	}
	return p
}
