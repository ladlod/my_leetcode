package common_list

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：

//  L0 → L1 → … → Ln-1 → Ln
// 请将其重新排列后变为：

// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …

// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/reorder-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：寻找中点，反转链表，之后双指针

func ReorderList(head *ListNode) {
	// step1: find mid
	quick, slow := head, head
	for quick != nil && quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}

	mid := slow.Next
	slow.Next = nil

	if mid == nil {
		return
	}
	head.Print()
	mid.Print()

	// step2: reverse
	pre := mid.Next
	mid.Next = nil
	for mid != nil && pre != nil {
		pNext := pre.Next
		pre.Next = mid
		mid = pre
		pre = pNext
	}

	head.Print()
	mid.Print()

	//step3: mix
	for p1, p2 := head, mid; p1 != nil && p2 != nil; {
		n1, n2 := p1.Next, p2.Next
		p1.Next, p2.Next = p2, p1.Next
		p1, p2 = n1, n2
	}

	head.Print()
}
