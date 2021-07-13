package dp

func rob213(nums []int) int {
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
		a := rob198(nums[1:])
		b := rob198(nums[:l-1])
		if a > b {
			return a
		} else {
			return b
		}
	}
}
