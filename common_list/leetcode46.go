package common_list

import "fmt"

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 思路：
// 递归回溯
func Permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i, num := range nums {
		v1 := make([]int, i)
		v2 := make([]int, len(nums)-i-1)
		copy(v1, nums[:i])
		copy(v2, nums[i+1:])
		next := Permute(append(v1, v2...))
		fmt.Println(next)
		for _, n := range next {
			res = append(res, append(n, num))
		}
	}

	return res
}
