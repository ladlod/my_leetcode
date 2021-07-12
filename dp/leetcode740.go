package main

func deleteAndEarn740(nums []int) int {
	numMap := make(map[int]int)
	max := 0
	for _, num := range nums {
		numMap[num]++
		if num > max {
			max = num
		}
	}
	if max <= 0 {
		return 0
	} else if max == 1 {
		return numMap[1]
	} else if max == 2 {
		if numMap[1] > numMap[2]*2 {
			return numMap[1]
		} else {
			return numMap[2] * 2
		}
	} else {
		res := make([]int, max)
		res[0] = numMap[1]
		if numMap[1] > numMap[2]*2 {
			res[1] = numMap[1]
		} else {
			res[1] = numMap[2] * 2
		}
		for i := 2; i < max; i++ {
			if numMap[i+1]*(i+1)+res[i-2] > res[i-1] {
				res[i] = numMap[i+1]*(i+1) + res[i-2]
			} else {
				res[i] = res[i-1]
			}
		}
		return res[max-1]
	}
}
