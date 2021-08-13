package common_list

// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

// 说明：不允许修改给定的链表。

// 进阶：

// 你是否可以使用 O(1) 空间解决此题？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/linked-list-cycle-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：快慢指针遍历
// 设入环位置为k，相遇位置为p，环长为l，快指针走了n+1圈，慢指针走了n圈
// 则 快指针共走了 k + p + (n+1) * l 慢指针共走了 k + p + n * l
// 因为快指针走了慢指针移动距离的二倍减1 所以 k + p + (n+1) * l = 2（k + p + n * l） - 1
// 可得 (1-n)l = k + p + 1,即 慢指针出发了环长的整数倍减一
// 所以，此时从原点出发一个指针，与慢指针一起前进，二者相遇的节点即为环的起点

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	quick := head.Next
	for quick != slow {
		if quick == nil || quick.Next == nil {
			return nil
		}
		slow = slow.Next
		quick = quick.Next.Next
	}

	flag := head
	slow = slow.Next
	for flag != slow {
		slow = slow.Next
		flag = flag.Next
	}

	return flag
}
