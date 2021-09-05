package week_contest

// 你需要访问 n 个房间，房间从 0 到 n - 1 编号。同时，每一天都有一个日期编号，从 0 开始，依天数递增。你每天都会访问一个房间。

// 最开始的第 0 天，你访问 0 号房间。给你一个长度为 n 且 下标从 0 开始 的数组 nextVisit 。在接下来的几天中，你访问房间的 次序 将根据下面的 规则 决定：

// 假设某一天，你访问 i 号房间。
// 如果算上本次访问，访问 i 号房间的次数为 奇数 ，那么 第二天 需要访问 nextVisit[i] 所指定的房间，其中 0 <= nextVisit[i] <= i 。
// 如果算上本次访问，访问 i 号房间的次数为 偶数 ，那么 第二天 需要访问 (i + 1) mod n 号房间。
// 请返回你访问完所有房间的第一天的日期编号。题目数据保证总是存在这样的一天。由于答案可能很大，返回对 109 + 7 取余后的结果。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/first-day-where-you-have-been-in-all-the-rooms
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// dp
// 由于nextVisit[i] <= i, 所以每次第一次到达i号房间时，必定要回到j=nextVisit[i]号房间，想要到达i+1号房间，则需要重新从j走到i
// 当访问到i+1号房间时，j->i号房间必定走过偶数次，可以推算出在从j回到i的过程中，j到i-1的房间也必定走过偶数次
// 我们可以记录，第1次走到房间i所需的天数为dp[i]，那么
// dp[i] = dp[i-1] (从0到dp[i-1]需要走的次数)
//		   + 1 (从dp[i-1]向前一步走到dp[i]的次数)
//		   + 1(从dp[i-1]回到dp[j]走的次数)
//		   + (dp[i-1] - dp[j])(从dp[j]回到dp[i-1]走的次数)

func FirstDayBeenInAllRooms(nextVisit []int) int {
	const mod = 1e9 + 7

	dp := make([]int, len(nextVisit))
	for i := 0; i < len(nextVisit); i++ {
		if i == 0 {
			dp[i] = 1
		} else {
			dp[i] = (2*dp[i-1] - dp[nextVisit[i-1]] + 2 + mod) % mod
		}
	}
	// fmt.Println(dp)
	return (dp[len(dp)-1] + mod - 1) % mod
}
