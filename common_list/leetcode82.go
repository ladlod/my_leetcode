package common_list

// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。

// 返回同样按升序排列的结果链表。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 顺序遍历链表，记录当前遍历状态
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := &ListNode{
		Next: head,
	}
	flag := false
	for p := pre; p.Next != nil; {
		if p.Next.Next != nil && p.Next.Val == p.Next.Next.Val {
			p.Next = p.Next.Next
			flag = true
		} else if flag {
			p.Next = p.Next.Next
			flag = false
		} else {
			p = p.Next
		}
	}

	return pre.Next
}
