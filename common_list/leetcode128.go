package common_list

import "fmt"

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

// eg:：

// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
func LongestConsecutive(nums []int) int {
	numM := make(map[int]bool)

	// 构造map
	for _, num := range nums {
		numM[num] = true
	}

	res := make(map[int]int)
	for _, num := range nums {
		if res[num] > 0 {
			continue
		}
		for k := 0; numM[num-k]; k++ {
			res[num-k] = k + 1
		}
	}

	// [1 2 0 1]
	// {}

	fmt.Println(res)

	r := 0
	for _, v := range res {
		if v > r {
			r = v
		}
	}

	return r
}
