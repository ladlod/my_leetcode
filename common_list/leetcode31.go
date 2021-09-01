package common_list

// 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）。

// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

// 必须 原地 修改，只允许使用额外常数空间。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/next-permutation
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 从后向前遍历寻找第一个降序的位置，将第一个数字之后的后续数字升序排列（反转），调换第一个数字和后面第一个比它大的数字的位置

func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	reverseArray(nums, i, len(nums)-1)
	if i > 0 {
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}

func reverseArray(nums []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// 部分快排
func quickSortPart(nums []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		for nums[j] > nums[mid] && j > mid {
			j--
		}
		for nums[i] < nums[mid] && i < mid {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	quickSortPart(nums, start, mid)
	quickSortPart(nums, mid+1, end)
}
