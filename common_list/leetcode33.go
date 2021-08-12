package common_list

// 整数数组 nums 按升序排列，数组中的值 互不相同 。

// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/*
思路：
解法1，遍历暴力求解
解法2，二分法 step1，二分法找到旋转处，step2，判断数字与旋转节点大小关系，二分查找节点左侧或右侧

样例：
[4,5,6,7,0,1,2]
0
*/
func Search(nums []int, target int) int {
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	} else {
		flag := searchFlag(nums, 0, len(nums)-1)
		if target > nums[len(nums)-1] {
			return searchResult(nums, target, 0, flag)
		} else {
			return searchResult(nums, target, flag+1, len(nums)-1)
		}
	}
}

func searchFlag(nums []int, left, right int) int {
	if left >= right {
		return -1
	}
	flag := (left + right) / 2
	if nums[flag] > nums[flag+1] {
		return flag
	} else if nums[flag] < nums[len(nums)-1] {
		return searchFlag(nums, left, flag)
	} else if nums[flag] > nums[len(nums)-1] {
		return searchFlag(nums, flag+1, right)
	}
	return -1
}

func searchResult(nums []int, target int, left, right int) int {
	if left > right {
		return -1
	} else if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}

	mid := (left + right) / 2
	if target == nums[mid] {
		return mid
	} else if target < nums[mid] {
		return searchResult(nums, target, left, mid)
	} else {
		return searchResult(nums, target, mid+1, right)
	}
}
