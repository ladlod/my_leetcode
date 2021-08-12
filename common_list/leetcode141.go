package common_list

// 给定一个链表，判断链表中是否有环。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

// 如果链表中存在环，则返回 true 。 否则，返回 false 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/linked-list-cycle
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/*
思路：
解法1：暴力 map标记
解法2：快慢指针，一个快指针，每次走两位，一个慢指针，每次走一位，遍历过程若追赶过程出现指针相同，则存在环
*/
func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)
	for p := head; p != nil; p = p.Next {
		if _, ok := nodeMap[p]; ok {
			return true
		} else {
			nodeMap[p] = true
		}
	}

	return false
}

func hasCycleV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	pre, next := head, head.Next
	for pre != next {
		if next == nil || next.Next == nil {
			return false
		}
		pre = pre.Next
		next = next.Next.Next
	}
	return true
}
