package common_list

// 统计所有小于非负整数 n 的质数的数量。

/*
思路：
解法1：暴力法，逐个遍历
解法2：任意一个数的倍数都不是质数，把所有质数的倍数标记为非质数，统计质数
解法3：
*/

func CountPrimesV1(n int) int {
	if n <= 1 {
		return 0
	}
	res := 1
	for i := 1; i < n; i++ {
		if i&1 == 0 && i != 2 {
			res++
			continue
		}
		for j := 3; j*j <= i; j += 2 { // 已经跳过了所有偶数，直接遍历奇数
			if i%j == 0 {
				res++
				break
			}
		}
	}
	return n - 1 - res
}

func CountPrimesV2(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	res := make([]int, n)
	res[1] = 1
	ne := 1
	for i := 2; i < n; i++ {
		if res[i] == 0 {
			for j := 2; j*i < n; j++ {
				if res[j*i] == 0 {
					res[j*i] = 1
					ne++
				}
			}
		}
	}

	return n - 1 - ne
}
