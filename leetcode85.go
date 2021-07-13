package main

func maximalRectangle(matrix [][]byte) int {
	h := len(matrix)
	if h <= 0 {
		return 0
	}
	w := len(matrix[0])

	heightP := make([][]int, h)
	for j := 0; j < h; j++ {
		heightP[j] = make([]int, w)
		for i := 0; i < w; i++ {
			if j == 0 {
				heightP[j][i] = int(matrix[j][i])
			} else {
				if matrix[j][i] == 1 {
					heightP[j][i] = heightP[j-1][i] + 1
				}
			}
		}
	}

	res := 0
	for j := 0; j < h; j++ {
		resP := 0
		for i := 0; i < w; i++ {
			if heightP[j][i] == 0 {
				continue
			}
			width := 1
			for l := i - 1; l >= 0 && heightP[j][l] >= heightP[j][i]; l-- {
				width++
			}
			for r := i + 1; r < w && heightP[j][r] >= heightP[j][i]; r++ {
				width++
			}
			if width*heightP[j][i] >= resP {
				resP = width * heightP[j][i]
			}
		}
		if resP >= res {
			res = resP
		}
	}
	return res
}
