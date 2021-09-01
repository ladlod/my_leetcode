package common_list

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：双指针法
// 设链表A，B重合长度为c，链表A长度为a+c, 链表B长度为b+c
// 可令两个指针从两个链表投开始，均行进a+b+c的路程相遇，相遇点即为链表交叉起始
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for pa, pb := headA, headB; ; {
		if pa == pb {
			return pa
		}
		if pa.Next == nil && pb.Next == nil {
			return nil
		}
		if pa.Next == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb.Next == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
}
