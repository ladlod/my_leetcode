package common_list

// 城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回由这些建筑物形成的 天际线 。

// 每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti] 表示：

// lefti 是第 i 座建筑物左边缘的 x 坐标。
// righti 是第 i 座建筑物右边缘的 x 坐标。
// heighti 是第 i 座建筑物的高度。
// 天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0 ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

// 注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/the-skyline-problem
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type sortMultiNums [][]int

func (s sortMultiNums) Len() int {
	return len(s)
}

func (s sortMultiNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortMultiNums) Less(i, j int) bool {
	if len(s[i]) < 3 || len(s[j]) < 3 {
		return false
	}
	if s[i][0] == s[j][0] {
		return s[i][2] < s[j][2]
	}
	return s[i][0] < s[j][0]
}

func GetSkyline(buildings [][]int) [][]int {
	res := make([][]int, 0)

	return res
}
