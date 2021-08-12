package interview

// 给定两个用链表表示的整数，每个节点包含一个数位。
// 这些数位是反向存放的，也就是个位排在链表首部
// 编写函数对这两个整数求和，并用链表形式返回结果。

/*
思路：从头遍历两个链表，求和并记录进位
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	inc := 0
	res := new(ListNode)
	p1, p2 := l1, l2
	r := res
	for p1 != nil || p2 != nil {
		v1, v2 := 0, 0
		if p1 != nil {
			v1 = p1.Val
		}
		if p2 != nil {
			v2 = p2.Val
		}
		if v1+v2+inc < 10 {
			r.Val = v1 + v2 + inc
			inc = 0
		} else {
			r.Val = (v1 + v2 + inc) % 10
			inc = (v1 + v2 + inc) / 10
		}
		if (p1 != nil && p1.Next != nil) || (p2 != nil && p2.Next != nil) {
			r.Next = new(ListNode)
			r = r.Next
		}
		if p1 != nil {
			p1 = p1.Next
		}
		if p2 != nil {
			p2 = p2.Next
		}
	}
	if inc > 0 {
		r.Next = &ListNode{
			Val: inc,
		}
	}

	return res
}
