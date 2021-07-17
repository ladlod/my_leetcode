package daily

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	d := len(nums) / 2
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i] < nums[d] {
			i++
		} else if nums[j] > nums[d] {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	QuickSort(nums[:d])
	QuickSort(nums[d+1:])
}
