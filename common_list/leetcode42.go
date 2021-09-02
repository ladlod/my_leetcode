package common_list

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 思路：双指针法
// 记录左右最高值，左值比右值小时，正序遍历，加上空余面积，右值比左值小时，逆序遍历
func Trap(height []int) int {
	res := 0
	lmax, rmax := 0, 0
	for l, r := 0, len(height)-1; l <= r; {
		if height[l] < height[r] {
			if height[l] > lmax {
				lmax = height[l]
			} else {
				res += lmax - height[l]
			}
			l++
		} else {
			if height[r] > rmax {
				rmax = height[r]
			} else {
				res += rmax - height[r]
			}
			r--
		}
	}
	return res
}
