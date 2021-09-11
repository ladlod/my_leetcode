package common_list

import (
	"strconv"
)

// 给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

// 示例 1 :
// 输入: 2736
// 输出: 7236
// 解释: 交换数字2和数字7。

// 示例 2 :
// 输入: 9973
// 输出: 9973
// 解释: 不需要交换。

// 给定数字的范围是 [0, 1e8]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-swap
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 任意一组数字，逆序排列的组合最大，所以最大数字组合从后向前遍历应该为顺序
// 我们可以记录每个数字最后出现的下标，从左到右扫描数字时，调换后面最大数字与当前数字的位置（若成功调换则返回，没调换则继续向后遍历）
func MaximumSwap(num int) int {
	strNum := []byte(strconv.Itoa(num))
	pos := make([]int, 10)
	for i := 0; i < len(strNum); i++ {
		pos[int(strNum[i]-'0')] = i
	}

	for i := 0; i < len(strNum); i++ {
		for j := 9; j > int(strNum[i]-'0'); j-- {
			if pos[j] > i {
				strNum[i], strNum[pos[j]] = strNum[pos[j]], strNum[i]
				v, _ := strconv.Atoi(string(strNum))
				return v
			}
		}
	}

	return num
}
