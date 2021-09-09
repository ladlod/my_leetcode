package daily

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	base := nums[0]
	l, r := 0, len(nums)-1

	for l < r {
		for l < r && nums[r] >= base {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= base {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = base
	QuickSort(nums[:l])
	QuickSort(nums[l+1:])
}
