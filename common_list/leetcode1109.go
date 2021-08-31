package common_list

// 这里有 n 个航班，它们分别从 1 到 n 进行编号。

// 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

// 请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。

//

// 示例 1：
// 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
// 输出：[10,55,45,25,25]
// 解释：
// 航班编号        1   2   3   4   5
// 预订记录 1 ：   10  10
// 预订记录 2 ：       20  20
// 预订记录 3 ：       25  25  25  25
// 总座位数：      10  55  45  25  25
// 因此，answer = [10,55,45,25,25]

// 示例 2：
// 输入：bookings = [[1,2,10],[2,2,15]], n = 2
// 输出：[10,25]
// 解释：
// 航班编号        1   2
// 预订记录 1 ：   10  10
// 预订记录 2 ：       15
// 总座位数：      10  25
// 因此，answer = [10,25]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/corporate-flight-bookings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, booking := range bookings {
		if len(booking) < 3 {
			continue
		}
		for i := booking[0] - 1; i < booking[1]; i++ {
			if i >= n {
				continue
			}
			res[i] += booking[2]
		}
	}

	return res
}

// 优化思路：
// 查分数组：如[1,2,3,4]为一个数组，其查分数组为[1,1,1,1]，对区间[1,2]+1，即对差分数组的第一位+1，第三位-1
//         将差分数组转化为[1,2,1,0]，转化为原值为[1,3,4,4]，可以此思路将问题转化为差分数组

func corpFlightBookingsV2(bookings [][]int, n int) []int {
	subR := make([]int, n)
	for _, bookings := range bookings {
		if len(bookings) < 3 {
			continue
		}
		subR[bookings[0]-1] += bookings[2]
		if bookings[1] < n {
			subR[bookings[1]] -= bookings[2]
		}
	}

	for i := 1; i < n; i++ {
		subR[i] = subR[i] + subR[i-1]
	}

	return subR
}
