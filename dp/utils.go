package dp

import "math"

func min(args ...int) int {
	res := math.MaxInt32
	for _, i := range args {
		if i < res {
			res = i
		}
	}
	return res
}

func max(args ...int) int {
	res := math.MinInt32
	for _, i := range args {
		if i > res {
			res = i
		}
	}
	return res
}
