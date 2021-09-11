package common_list

import "fmt"

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数 。
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

// 示例 1：
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2

// 示例 2：
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

// 示例 3：
// 输入：nums1 = [0,0], nums2 = [0,0]
// 输出：0.00000
// 示例 4：
// 输入：nums1 = [], nums2 = [1]
// 输出：1.00000

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 常规思路
// 统计两个数组长度和，求出中点位置，归并排序两个数组
// 进阶思路
// 二分查找
// 找nums1和nums2的中位数，可以理解为找到第l/2位置的数和第l/2+1位置的数

// 问题转换成找到两个有序数组中的第n个数
// 找到nums1的中点和nums2的中点，比较二者大小
// 找到两个数组中的第k/2-1个数，比较大小
// 若nums1[k/2-1] < nums2[k/2-1]说明在nums1[k/2-1]之前，最多有k-2个数，不可能存在第k个数，所以可以从nums1[k/2:]和nums2中继续寻找第(k-k/2+1)个数
// 反之nums1[k/2-1] > nums2[k/2-1]，则在nums2[k/2:]和nums1中继续寻找第(k-k/2+1)个数
// 若nums1[k/2-1] = nums2[k/2-1], 则说明在该数之前有确定的k-2个数，答案不可能在该数之前，所以可以从nums1[k/2:]和nums2[k/2:]中寻找第(k-k/2-k/2+2)个数
// 边界条件若遇到k/2-1值溢出则需特殊考虑
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l&1 == 0 {
		return (float64(findKthSortedArrays(nums1, nums2, l/2)) + float64(findKthSortedArrays(nums1, nums2, l/2+1))) / 2
	} else {
		return float64(findKthSortedArrays(nums1, nums2, (l+1)/2))
	}
}

func findKthSortedArrays(nums1 []int, nums2 []int, k int) int {
	fmt.Println(nums1, nums2, k)
	if len(nums1) == 0 {
		return nums2[k-1]
	} else if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	k1, k2 := max(1, min((k+1)/2-1, len(nums1))), max(1, min((k+1)/2-1, len(nums2)))

	if nums1[k1-1] < nums2[k2-1] {
		return findKthSortedArrays(nums1[k1:], nums2, k-k1)
	} else if nums1[k1-1] > nums2[k2-1] {
		return findKthSortedArrays(nums1, nums2[k2:], k-k2)
	} else {
		if k1+k2 == k {
			return nums1[k1-1]
		} else {
			return findKthSortedArrays(nums1[k1:], nums2[k2:], k-k1-k2)
		}
	}
}
