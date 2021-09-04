package common_list

import "fmt"

// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// eg:[6,1,3,2,4,7] -> [0,-5,2,-1,2,3]
// 思路：求一段最大收益，即可理解为求差分数组最大和
//      同理，求两段最大收益，可理解为求差分数组中两段的最大和=>差分数组某一区间的和，减去这一区间的最小值（可以不减）
// 时间复杂度 O(n^2)

// 思路2：求两段的最大收益可以理解为求某一点左侧最大收益和某一点右侧最大收益的和
// 先一次循环遍历记录到某点i为止的最大收益数组 []int maxL
// 再计算，从i到结束点的最大收益求和
func MaxProfit3(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	maxL := make([]int, len(prices))
	min := prices[0]
	maxL[0] = 0

	// 计算到i为止，左侧最大值
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > maxL[i-1] {
			maxL[i] = prices[i] - min
		} else {
			maxL[i] = maxL[i-1]
		}
	}

	// 计算从i开始，右侧最大值
	maxR := make([]int, len(prices))
	max := prices[len(prices)-1]
	maxR[len(prices)-1] = 0
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > max {
			max = prices[i]
		}
		if max-prices[i] > maxR[i+1] {
			maxR[i] = max - prices[i]
		} else {
			maxR[i] = maxR[i+1]
		}
	}

	fmt.Println(maxL)
	fmt.Println(maxR)

	res := 0
	for i := 0; i < len(prices); i++ {
		if maxL[i]+maxR[i] > res {
			res = maxL[i] + maxR[i]
		}
	}

	return res
}
