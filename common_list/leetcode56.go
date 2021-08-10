package common_list

import "sort"

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-intervals
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
思路:
记录当前最做，最右数值，遍历
最右侧比当前小，跳过
最右侧比当前大，比较左侧与当前最右侧,若小于当前最右侧,更新最右侧数值
                              ,若大于当前最右侧,当前最左最右组成的数组入队,更新最左最右值
*/
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)

	if len(intervals) <= 0 {
		return res
	}

	sort.Slice(intervals, func(i, j int) bool {
		if len(intervals[i]) < 1 {
			return true
		} else if len(intervals[j]) < 1 {
			return false
		}
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[1] <= right {
			continue
		} else {
			if interval[0] <= right {
				right = interval[1]
			} else {
				res = append(res, []int{left, right})
				left = interval[0]
				right = interval[1]
			}
		}
	}
	res = append(res, []int{left, right})
	return res
}
