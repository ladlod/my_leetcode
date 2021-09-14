package common_list

import "math"

// 给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

// 如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/132-pattern
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 暴力法，O(n^2)，超时
func find132patternV1(nums []int) bool {
	lMin := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		for k := i + 1; k < len(nums); k++ {
			if lMin < nums[k] && nums[k] < nums[i] {
				return true
			}
		}
		if nums[i] < lMin {
			lMin = nums[i]
		}
	}
	return false
}

// 单调栈，维护从右至左的最大值，记录所有出栈元素的值的最大值为次大值，只要保证遍历到的i比次大值小及存在132模式
func find132patternV2(nums []int) bool {
	stack := []int{}
	maxnd := math.MinInt32
	for j := len(nums) - 1; j >= 0; j-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums[j] {
			if stack[len(stack)-1] > maxnd {
				maxnd = stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[j])
		if nums[j] < maxnd {
			return true
		}
	}

	return false
}
