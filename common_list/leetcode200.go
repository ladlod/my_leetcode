package common_list

import "fmt"

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-islands
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 深度优先搜索
// 并查集

type pSet []int

func initPset(size int) pSet {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}
	return s
}

func (s pSet) find(i int) int {
	if s[i] != i {
		s[i] = s.find(s[i])
	}
	return s[i]
}

func (s pSet) merge(a, b int) {
	s[s.find(b)] = s[a]
}

func NumIslands(grid [][]byte) int {
	h := len(grid)
	if h <= 0 {
		return 0
	}
	l := len(grid[0])

	p := initPset(l * h)
	num0 := 0
	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			if grid[i][j] == '1' {
				if i > 0 && grid[i-1][j] == '1' {
					p.merge(i*l+j, (i-1)*l+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					p.merge(i*l+j, i*l+j-1)
				}
			} else {
				num0++
			}
		}
	}

	res := 0
	for i := 0; i < l*h; i++ {
		if p.find(i) == i {
			res++
		}
	}

	fmt.Println(p)

	return res - num0
}
