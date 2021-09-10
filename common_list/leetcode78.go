package common_list

import (
	"fmt"
)

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：

// 输入：nums = [0]
// 输出：[[],[0]]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/subsets
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
func Subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})

	for _, num := range nums {
		l := len(res)
		for i := 0; i < l; i++ {
			res = append(res, append(res[i], num))
			fmt.Println(res, num)
		}
	}

	return res
}

// 此处应注意go切片的实现及扩容原理
// type slice struct {
// 	ptr unsafe.Pointer // Array pointer
// 	len int            // slice length
// 	cap int            // slice capacity
// }
