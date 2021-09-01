package common_list

// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

// 请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/odd-even-linked-list
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 双指针法
func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even, h2 := head, head.Next, head.Next
	for even != nil && even.Next != nil {
		odd.Next, even.Next = even.Next, odd.Next.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = h2

	return head
}
