package common_list

import "sort"

// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的每个数字在每个组合中只能使用一次。

// 注意：解集不能包含重复的组合。

// 1 <= target <= 30

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combination-sum-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// candidates = [10,1,2,7,6,1,5], target = 8
// sort [1,1,2,5,6,7,10]
// 回溯
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	return innerCombinationSum2(candidates, target)
}

func innerCombinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(candidates); i++ {
		if candidates[i] < target {
			for _, v := range innerCombinationSum2(candidates[i+1:], target-candidates[i]) {
				res = append(res, append([]int{candidates[i]}, v...))
			}
		} else if candidates[i] == target {
			res = append(res, []int{candidates[i]})
		}
		for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
			i++
		}
	}
	return res
}
