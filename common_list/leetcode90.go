package common_list

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	rtn := make([][]int, 0)
	rtn = append(rtn, []int{})

	sort.Sort(sort.IntSlice(nums))

	flag := 0
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			flag++
			continue
		} else {
			if flag == 0 {
				for _, tmp := range rtn {
					t := make([]int, len(tmp))
					copy(t, tmp)
					rtn = append(rtn, append(t, nums[i]))
				}
			} else {
				tmp2 := [][]int{[]int{nums[i]}}
				for ; flag > 0; flag-- {
					tmp2 = append(tmp2, append(tmp2[len(tmp2)-1], nums[i]))
				}
				for _, tmp := range rtn {
					for _, tmp3 := range tmp2 {
						t := make([]int, len(tmp))
						copy(t, tmp)
						rtn = append(rtn, append(t, tmp3...))
					}
				}
				flag = 0
			}
		}

	}

	return rtn
}
