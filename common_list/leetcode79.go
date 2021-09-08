package common_list

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-search
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 从网格中所有字符串的第一个字符开始进行深度优先搜索
func ExistStr(board [][]byte, word string) bool {
	if len(word) <= 0 {
		return true
	}
	h := len(board)
	if h <= 0 {
		return false
	}
	l := len(board[0])
	found := make([][]bool, h)
	for i := 0; i < h; i++ {
		found[i] = make([]bool, l)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < l; j++ {
			if dfs(board, word, found, i, j) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, found [][]bool, i, j int) bool {
	if len(word) <= 0 {
		return true
	} else if len(word) == 1 {
		return board[i][j] == word[0]
	}

	if board[i][j] != word[0] {
		return false
	}
	found[i][j] = true
	if i > 0 && !found[i-1][j] {
		if dfs(board, word[1:], found, i-1, j) {
			return true
		}
	}
	if j > 0 && !found[i][j-1] {
		if dfs(board, word[1:], found, i, j-1) {
			return true
		}
	}
	if i < len(board)-1 && !found[i+1][j] {
		if dfs(board, word[1:], found, i+1, j) {
			return true
		}
	}
	if j < len(board[0])-1 && !found[i][j+1] {
		if dfs(board, word[1:], found, i, j+1) {
			return true
		}
	}
	found[i][j] = false
	return false
}
