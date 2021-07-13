package dp

func rob198(nums []int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[1] > nums[0] {
			return nums[1]
		} else {
			return nums[0]
		}
	} else {
		res := make([]int, l)
		res[0] = nums[0]
		if nums[1] > nums[0] {
			res[1] = nums[1]
		} else {
			res[1] = nums[0]
		}
		for i := 2; i < l; i++ {
			if res[i-2]+nums[i] >= res[i-1] {
				res[i] = res[i-2] + nums[i]
			} else {
				res[i] = res[i-1]
			}
		}
		return res[l-1]
	}
}
