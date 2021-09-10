package stab_offer

// 统计一个数字在排序数组中出现的次数。

// 示例 1:

// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: 2
// 示例 2:

// 输入: nums = [5,7,7,8,8,10], target = 6
// 输出: 0

// 二分法找到元素起止位置
func search(nums []int, target int) int {
	l := searchStart(nums, target, 0, len(nums)-1)
	if l < 0 {
		return 0
	}
	r := searchEnd(nums, target, l, len(nums)-1)
	// fmt.Println(l, r)
	return r - l + 1
}

func searchStart(nums []int, target int, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if (mid < 1 || nums[mid-1] < target) && nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return searchStart(nums, target, mid+1, r)
	} else if mid > 0 && nums[mid-1] >= target {
		return searchStart(nums, target, l, mid-1)
	}
	return -1
}

func searchEnd(nums []int, target int, l, r int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if nums[mid] == target && (mid >= len(nums)-1 || nums[mid+1] > target) {
		return mid
	} else if mid < len(nums)-1 && nums[mid+1] <= target {
		return searchEnd(nums, target, mid+1, r)
	} else if nums[mid] > target {
		return searchEnd(nums, target, l, mid-1)
	}
	return -1
}
