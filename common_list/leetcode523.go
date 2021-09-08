package common_list

// 给你一个整数数组 nums 和一个整数 k ，编写一个函数来判断该数组是否含有同时满足下述条件的连续子数组：

// 子数组大小 至少为 2 ，且
// 子数组元素总和为 k 的倍数。
// 如果存在，返回 true ；否则，返回 false 。

// 如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终视为 k 的一个倍数。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/continuous-subarray-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 暴力解法，枚举左右子数组，时间复杂度O(n^2)
func checkSubarraySum(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum%k == 0 {
				return true
			}
		}
	}

	return false
}

// 记录以i位置可以存在的余数
// dp[i] = sum(dp[i-1], nums[i] % k)
// 时间复杂度 O(kn)
func checkSubarraySumV2(nums []int, k int) bool {
	if len(nums) <= 0 {
		return false
	}

	subMap := make([]map[int]bool, len(nums))
	subMap[0] = map[int]bool{
		nums[0]: true,
	}
	for i := 1; i < len(nums); i++ {
		subMap[i] = map[int]bool{
			nums[i]: true,
		}
		for j := range subMap[i-1] {
			if (j+nums[i])%k == 0 {
				return true
			} else {
				subMap[i][(j+nums[i])%k] = true
			}
		}
	}

	return false
}

// 对于一组数，从0-i和为sumi，从0-j和为sumj，那么从i-j的和为sumj-sumi
// 所以，我们可以从0开始计算和，标记每个k的余数第一次出现的位置，如果同一个余数第二次出现，表示这两个位置中间的和是k的倍数
func checkSubarraySumV3(nums []int, k int) bool {
	subMap := map[int]int{
		0: -1,
	}
	sum := 0
	for i, num := range nums {
		sum += num
		if j, ok := subMap[sum%k]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			subMap[sum%k] = i
		}
	}
	return false
}
