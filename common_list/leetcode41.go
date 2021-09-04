package common_list

// 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

// 思路：
// 思路1 bitmap：32进制正整数最大值2^31-1，需要2^31个位来存储bitmap，即2^26个无符号32位正整数，可以从后向前bitmap判断第一个不存在的正整数
// 思路2 map变形：因为一个长度为n的数组，第一个没出现的正整数最大可能为n+1，所以只需要一个长度为n的map即可存储所有可能没有出现的正整数
//				 考虑题中要求使用常数级的空间解决，我们可以把数组当做一个map来使用，第i为存储出现过的i
func FirstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
