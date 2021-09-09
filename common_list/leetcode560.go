package common_list

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。

// 示例 1：
// 输入：nums = [1,1,1], k = 2
// 输出：2

// 示例 2：
// 输入：nums = [1,2,3], k = 3
// 输出：2

// 1 <= nums.length <= 2 * 10e4
// -1000 <= nums[i] <= 1000
// 10e-7 <= k <= 10e7

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 暴力解法：时间复杂度O(n^2)
func subarraySum(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}

	return res
}

// 前缀和
// 记录从0-n所有nums求和出现过的前缀和
// sum(i->k) = sum(k) - sum(i)
func subarraySumV2(nums []int, k int) int {
	sum, res := 0, 0
	mp := map[int]int{
		0: 1,
	}
	for _, num := range nums {
		sum += num
		if _, ok := mp[sum-k]; ok {
			res += mp[sum-k]
		}
		mp[sum]++
	}
	return res
}
