package common_list

// 实现 int sqrt(int x) 函数。

// 计算并返回 x 的平方根，其中 x 是非负整数。

// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/sqrtx
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func MySqrt(x int) int {
	if x <= 1 {
		return x
	}
	return mySqrtInner(0, x, x)
}

func mySqrtInner(l, h, x int) int {
	mid := (l + h) / 2
	// 如果溢出了int32位怎么办？应改为 mid <= x /mid && mid+1 > x/(mid+1)
	// go的int位数和go版本有关，当前默认为64位，所以没影响
	if mid*mid <= x && (mid+1)*(mid+1) > x {
		return mid
	} else if mid*mid > x {
		return mySqrtInner(l, mid, x)
	} else {
		return mySqrtInner(mid+1, h, x)
	}
}
