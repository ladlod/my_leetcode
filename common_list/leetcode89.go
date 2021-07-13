package common_list

func grayCode(n int) []int {
	rtn := make([]int, 0)
	rtn = append(rtn, 0)

	nums := make([]int, n)

	_map := make(map[int]bool)
	_map[0] = true

tag:
	for true {
		for i := 0; i < n; i++ {
			if nums[i] == 0 {
				nums[i] = 1
			} else {
				nums[i] = 0
			}
			num := arrayToNum(nums)
			if _, ok := _map[num]; !ok {
				rtn = append(rtn, num)
				_map[num] = true
				continue tag
			}
			if nums[i] == 0 {
				nums[i] = 1
			} else {
				nums[i] = 0
			}
		}
		break
	}

	return rtn
}

func arrayToNum(num []int) (rtn int) {
	for i := 0; i < len(num); i++ {
		rtn = rtn*2 + num[i]
	}
	return
}
