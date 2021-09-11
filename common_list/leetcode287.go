package common_list

import "fmt"

// 给定一个包含 n + 1 个整数的数组 nums ，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。

// 假设 nums 只有 一个重复的整数 ，找出 这个重复的数 。

// 你设计的解决方案必须不修改数组 nums 且只用常量级 O(1) 的额外空间。

// 示例 1：

// 输入：nums = [1,3,4,2,2]
// 输出：2
// 示例 2：

// 输入：nums = [3,1,3,4,2]
// 输出：3
// 示例 3：

// 输入：nums = [1,1]
// 输出：1
// 示例 4：

// 输入：nums = [1,1,2]
// 输出：1

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-the-duplicate-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// bitmap
func findDuplicateV1(nums []int) int {
	bitmap := make([]int, (len(nums)-1)/32+1)
	for _, num := range nums {
		if bitmap[num/32]&(1<<((num-1)%32)) > 0 {
			return num
		} else {
			bitmap[num/32] |= 1 << ((num - 1) % 32)
		}
		// fmt.Println(bitmap)
	}

	return -1
}

// 思路
// 把数组看成一个链表，快慢指针
// 由于数组长度n+1，里面可能的数字范围为1-n，所以我们可以把数组的0号位看做图的起点，链表的边由i -> nums[i], 因为数组中存在相同数字，所以链表必定存在至少一个环
// 且环的入口即为重复数字

// 寻找带环链表环的入口
// 快慢指针同时出发，快指针行动速度为慢指针2倍
// 设环长为n，相遇位置距入环口为a，入环位置为b，则b+a = kn, k为正整数
// 所以在快慢指针相遇点，同时从原点出发一个慢指针，两个慢指针相遇的点即为入环点

func FindDuplicateV2(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for ; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}

	fmt.Println(slow, fast)

	res := 0
	for ; res != slow; res, slow = nums[res], nums[slow] {
	}

	return res
}
