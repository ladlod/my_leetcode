package week_contest

import "fmt"

// 给你一个 下标从 0 开始 的整数数组 nums ，返回满足下述条件的 不同 四元组 (a, b, c, d) 的 数目 ：

// nums[a] + nums[b] + nums[c] == nums[d] ，且
// a < b < c < d

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/count-special-quadruplets
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 暴力解法
func countQuadruplets(nums []int) int {
	res := 0
	for j := len(nums) - 1; j > 2; j-- {
		for i := 0; i < j-2; i++ {
			for p := i + 1; p < j-1; p++ {
				for q := p + 1; q < j; q++ {
					if nums[i]+nums[p]+nums[q] == nums[j] {
						fmt.Println(i, p, q, j)
						res++
					}
				}
			}
		}
	}
	return res
}
