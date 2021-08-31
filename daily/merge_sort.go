package daily

import "fmt"

func MergeSort(nums []int) []int {
	fmt.Println(nums)
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	l := MergeSort(nums[:mid])
	r := MergeSort(nums[mid:])

	return merge(l, r)
}

func merge(nums1, nums2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	if i != len(nums1) {
		res = append(res, nums1[i:]...)
	}
	if j != len(nums2) {
		res = append(res, nums2[j:]...)
	}
	return res
}

func MergeSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	quick, slow := head, head
	for quick != nil && quick.Next != nil && quick.Next.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	if mid == nil {
		return head
	}

	return mergeList(MergeSortList(head), MergeSortList(mid))
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	pre := &ListNode{}
	p := pre
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			p.Next = head1
			head1 = head1.Next
		} else {
			p.Next = head2
			head2 = head2.Next
		}
		p = p.Next
	}
	if head1 != nil {
		p.Next = head1
	} else if head2 != nil {
		p.Next = head2
	}
	return pre.Next
}
