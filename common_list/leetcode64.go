package common_list

// 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。

// 思路：
// 动态规划
// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
func minPathSum(grid [][]int) int {
	h := len(grid)
	if h <= 0 {
		return 0
	}
	l := len(grid[0])

	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, l)
		for j := 0; j < l; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}

	return dp[h-1][l-1]
}
