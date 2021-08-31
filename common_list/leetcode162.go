package common_list

// 峰值元素是指其值大于左右相邻值的元素。

// 给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

// 你可以假设 nums[-1] = nums[n] = -∞ 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-peak-element
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findPeakElement(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	res := 0
	for i, num := range nums {
		if num >= nums[res] {
			res = i
		} else {
			break
		}
	}
	return res
}

// 二分法
func FindPeakElementV2(nums []int) int {
	return innerFindPeakElementV2(nums, 0, len(nums)-1)
}

func innerFindPeakElementV2(nums []int, l, r int) int {
	mid := (l + r) / 2

	if ((mid-1 >= 0 && nums[mid] > nums[mid-1]) || (mid-1 < 0)) &&
		((mid+1 < len(nums) && nums[mid] > nums[mid+1]) || mid+1 >= len(nums)) {
		return mid
	}
	if l < r {
		if v := innerFindPeakElementV2(nums, l, mid); v >= 0 {
			return v
		}
		if v := innerFindPeakElementV2(nums, mid+1, r); v >= 0 {
			return v
		}
	}

	return -1
}
