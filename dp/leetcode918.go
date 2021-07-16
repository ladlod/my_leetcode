package dp

func MaxSubarraySumCircular918(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	res := maxSubarraySumCircular918_1(nums, 1)

	if tmp := sum - maxSubarraySumCircular918_1(nums[1:], -1); tmp > res {
		res = tmp
	}

	if tmp := sum - maxSubarraySumCircular918_1(nums[:len(nums)-1], -1); tmp > res {
		res = tmp
	}

	return res
}

func maxSubarraySumCircular918_1(nums []int, sign int) int {
	l := len(nums)
	if l <= 0 {
		return 0
	}
	now := nums[0] * sign
	res := now
	for i := 1; i < l; i++ {
		if now > 0 {
			now += nums[i] * sign
		} else {
			now = nums[i] * sign
		}
		if now > res {
			res = now
		}
	}
	return res * sign
}
