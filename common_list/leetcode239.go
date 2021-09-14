package common_list

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

// 返回滑动窗口中的最大值。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sliding-window-maximum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 单调栈
func maxSlidingWindow(nums []int, k int) []int {
	type stackNode struct {
		pos int
		val int
	}

	stack := make([]*stackNode, 0)
	for i := 0; i < k; i++ {
		for len(stack) > 0 && stack[len(stack)-1].val < nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, &stackNode{
			pos: i,
			val: nums[i],
		})
	}

	res := []int{stack[0].val}
	for j := k; j < len(nums); j++ {
		if len(stack) > 0 && stack[0].pos+k <= j {
			stack = stack[1:]
		}
		for len(stack) > 0 && stack[len(stack)-1].val < nums[j] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, &stackNode{
			pos: j,
			val: nums[j],
		})
		res = append(res, stack[0].val)
	}

	return res
}
