package common_list

import (
	"fmt"
	"math"
)

// 思路：三次选择排序，选出第三大的值
func ThirdMax(nums []int) int {
	var res int = math.MaxInt32
	var max int = math.MinInt32

	k := 0
	for i := 0; i < 3; i++ {
		tmp := math.MinInt32
		for _, num := range nums {
			if num < res && num >= tmp {
				tmp = num
				k = i
			}
			if i == 0 && num > max {
				max = num
			}
		}
		if k == i {
			res = tmp
			fmt.Println(res)
		}
	}

	if k < 2 {
		res = max
	}
	return res
}
