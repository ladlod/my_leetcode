package common_list

// 珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

// 珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

// 返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/koko-eating-bananas
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 示例 1：

// 输入: piles = [3,6,7,11], H = 8
// 输出: 4
// 示例 2：

// 输入: piles = [30,11,23,4,20], H = 5
// 输出: 30

// 思路：
// 设吃香蕉的速度为k，那么 sum(piles[i]/k) <= h
// 暴力解法：从1->max(piles)，逐个尝试直到找到满足条件的k，时间复杂度O(mn), m为max(piles), n为len(piles)
// 暴力解法优化：二分查找
func minEatingSpeed(piles []int, h int) int {
	min := 1
	max := max(piles...)
	for min < max {
		mid := (min + max) / 2
		if !checkEatingSpeed(piles, h, mid) {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func checkEatingSpeed(piles []int, h, k int) bool {
	sum := 0
	for _, p := range piles {
		sum += (p-1)/k + 1
	}
	return sum <= h
}
