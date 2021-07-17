package common_list

import "sort"

func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	permuteUnique47(nums, res)

	return res
}

func permuteUnique47(nums []int, res [][]int) [][]int {
	for i := 0; i < len(nums); i++ {
		for _, _res := range res {
			_res = append(_res, nums[i])
		}
		permuteUnique47(append(nums[:i], nums[i+1:]...), res)
	}
	return res
}
