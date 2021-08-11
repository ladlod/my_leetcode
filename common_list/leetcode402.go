package common_list

import "fmt"

// 给你一个以字符串表示的非负整数 num 和一个整数 k ，移除这个数中的 k 位数字，使得剩下的数字最小。请你以字符串形式返回这个最小的数字。

/*
思路：
位数相同的数字高位越小则该数字越小
常规思路：从前向后，取1+k个数字，找到里面最小的数字，移除该数字前面的n个数字，k-n，继续向后遍历，直到k=0
*/

func RemoveKdigits(num string, k int) string {
	for i := 0; k > 0; i++ {
		if len(num)-i == k {
			num = num[:i]
			k = 0
		} else {
			n := i
			for j := i; j < i+k+1 && j < len(num); j++ {
				if num[j] < num[n] {
					n = j
				}
			}
			fmt.Println(i, k, n)
			num = num[:i] + num[n:]
			k -= (n - i)
		}
	}

	i := 0
	for ; i < len(num); i++ {
		if num[i] != '0' {
			break
		}
	}

	if len(num)-i > 0 {
		return num[i:]
	} else {
		return "0"
	}
}
