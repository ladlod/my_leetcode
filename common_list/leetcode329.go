package common_list

import (
	"sort"
)

// 给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

// 对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// eg：
// [9,9,4]
// [6,6,8]
// [2,1,1]
// res: 1->2->6->9 len(4)

// 动态规划：任意一个数，从它开始可以走的最长路径，是它周围所有比它大的数的路径+1的最大值
// 按照所有数字从大到小遍历
// 时间复杂度O(mn*log(mn)) -> 优化方向：改为从所有最大的节点开始进行仅小于本身节点的广度优先搜索，可以将时间复杂度降至O(mn)
// 空间复杂度O(mn)
func LongestIncreasingPath(matrix [][]int) int {
	if len(matrix) <= 0 {
		return 0
	}

	l := len(matrix)
	h := len(matrix[0])

	type pos struct {
		x int
		y int
	}
	posList := make([]*pos, 0)
	for i := range matrix {
		for j := range matrix[0] {
			posList = append(posList, &pos{x: i, y: j})
		}
	}

	sort.Slice(posList, func(i, j int) bool {
		return matrix[posList[i].x][posList[i].y] > matrix[posList[j].x][posList[j].y]
	})

	dp := make([][]int, l)
	for i := range matrix {
		dp[i] = make([]int, h)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	res := 0
	for _, p := range posList {
		x, y := p.x, p.y
		if x > 0 && matrix[x-1][y] > matrix[x][y] && (dp[x-1][y]+1 > dp[x][y]) {
			dp[x][y] = dp[x-1][y] + 1
		}
		if x < l-1 && matrix[x+1][y] > matrix[x][y] && (dp[x+1][y]+1 > dp[x][y]) {
			dp[x][y] = dp[x+1][y] + 1
		}
		if y > 0 && matrix[x][y-1] > matrix[x][y] && (dp[x][y-1]+1 > dp[x][y]) {
			dp[x][y] = dp[x][y-1] + 1
		}
		if y < h-1 && matrix[x][y+1] > matrix[x][y] && (dp[x][y+1]+1 > dp[x][y]) {
			dp[x][y] = dp[x][y+1] + 1
		}
		if dp[x][y] > res {
			res = dp[x][y]
		}
		// fmt.Println(dp)
	}

	return res
}
