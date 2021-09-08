package common_list

// 给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。

// 返回仅包含 1 的最长（连续）子数组的长度。

// 示例 1：

// 输入：A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：
// [1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-consecutive-ones-iii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 滑动窗口
func LongestOnes(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}
	l, r, res := 0, 0, 0
	for r < len(nums) {
		if nums[r] == 1 {
			if r-l+1 > res {
				res = r - l + 1
			}
			r++
		} else {
			if k > 0 {
				k--
				if r-l+1 > res {
					res = r - l + 1
				}
				r++
			} else {
				if nums[l] == 1 {
					l++
				} else {
					l++
					k++
				}
			}
		}
	}

	return res
}
