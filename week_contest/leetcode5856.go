// 你被安排了 n 个任务。任务需要花费的时间用长度为 n 的整数数组 tasks 表示，第 i 个任务需要花费 tasks[i] 小时完成。一个 工作时间段 中，你可以 至多 连续工作 sessionTime 个小时，然后休息一会儿。

// 你需要按照如下条件完成给定任务：

// 如果你在某一个时间段开始一个任务，你需要在 同一个 时间段完成它。
// 完成一个任务后，你可以 立马 开始一个新的任务。
// 你可以按 任意顺序 完成任务。
// 给你 tasks 和 sessionTime ，请你按照上述要求，返回完成所有任务所需要的 最少 数目的 工作时间段 。

// 测试数据保证 sessionTime 大于等于 tasks[i] 中的 最大值 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 用一个数组存储所有任务子集的总耗时，数组的index可用一个bitmap表示
// dp 遍历集合的所有子集，若子集和小于等于sessionTime，dp[i] = min(dp[i], dp[i^j] + 1)

package week_contest

func MinSessions(tasks []int, sessionTime int) int {
	l := 1 << len(tasks)
	sum := make([]int, l)

	// 计算子集
	for i := range tasks {
		for k := 0; k < 1<<i; k++ {
			sum[k|1<<i] = sum[k] + tasks[i]
		}
	}

	// dp
	dp := make([]int, l)
	for i := range dp {
		dp[i] = len(tasks)
	}
	dp[0] = 0

	for i := range dp {
		for j := i; j > 0; j = (j - 1) & i {
			if sum[j] <= sessionTime && dp[j^i]+1 < dp[i] {
				dp[i] = dp[j^i] + 1
			}
		}
	}

	return dp[l-1]
}
