package common_list

// 给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素.

// eg
// 输入:
// [
//  [ 1, 2, 3 ],
//  [ 4, 5, 6 ],
//  [ 7, 8, 9 ],
// ]

// 输出:  [1,2,4,7,5,3,6,8,9]

// 思路1：
// 左上 j++ -> i++, j--
// 右上 i++ -> i--, j++
// 左下 j++ -> i++, j--
// 右下 return
// 上边 j++ -> i++, j--
// 左边 i++ -> i--, j++
// 下边 j++ -> i++, j--
// 右边 i++ -> i--, j++
func FindDiagonalOrder(mat [][]int) []int {
	res := make([]int, 0)
	h := len(mat)
	if h <= 0 {
		return res
	}
	l := len(mat[0])

	fi, fj := -1, 1
	flag := false
	for i, j := 0, 0; i < h && j < l; {
		res = append(res, mat[i][j])
		if i == 0 || i == h-1 || j == 0 || j == l-1 {
			if !flag {
				if i == 0 && j == l-1 { // 右上角
					i++
				} else if j == 0 && i == h-1 { // 左下角
					j++
				} else if i == 0 || i == h-1 { // 上下边
					j++
				} else if j == 0 || j == l-1 { // 左右边
					i++
				}
				if h != 1 && l != 1 {
					flag = !flag
				}
				continue
			} else {
				fi, fj = -fi, -fj
				flag = !flag
			}
		}
		i, j = i+fi, j+fj
	}

	return res
}

// 思路2：
// 遍历所有对角线，之后分奇偶组合
func FindDiagonalOrderV2(mat [][]int) []int {
	h := len(mat)
	if h <= 0 {
		return []int{}
	}
	l := len(mat[0])

	res := make([]int, 0)
	for i, j := 0, 0; i < h; {
		d := make([]int, 0)
		for x, y := i, j; y >= 0 && x < h; x, y = x+1, y-1 {
			d = append(d, mat[x][y])
		}
		if (i+j)&1 == 1 {
			res = append(res, d...)
		} else {
			for j := len(d) - 1; j >= 0; j-- {
				res = append(res, d[j])
			}
		}
		if j == l-1 {
			i++
		} else {
			j++
		}
	}
	return res
}
