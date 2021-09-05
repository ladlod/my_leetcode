package week_contest

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 nums ，你可以在 nums 上执行下述操作 任意次 ：

// 如果 gcd(nums[i], nums[j]) > 1 ，交换 nums[i] 和 nums[j] 的位置。其中 gcd(nums[i], nums[j]) 是 nums[i] 和 nums[j] 的最大公因数。
// 如果能使用上述交换方式将 nums 按 非递减顺序 排列，返回 true ；否则，返回 false 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/gcd-sort-of-an-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 对于一组数a b c，若gcd(a,b) > 1且gcd(b,c) > 1，则a b c可以随意交换顺序
// 一个数字和它的所有质因数放到一个集合里，对所有差集>0的集合求并集
// 对原数组进行排序，对两个数组同时遍历，遇到不相同数字时判断两个数字是否在同一个集合里，若在，则表示这两个数字可以调换位置，否则返回false

var o int = 0

// 并查集
type pSet []int

// 初始化并查集，所有元素单独属于一个集合
func initPSet(l int) pSet {
	res := make([]int, l+1)
	for i := 0; i <= l; i++ {
		res[i] = i
	}
	return res
}

// 查找元素所属集合
func (p pSet) find(v int) int {
	if p[v] != v {
		o++
		p[v] = p.find(p[v])
	}
	return p[v]
}

// 合并b和a所在的集合
func (p pSet) merge(a, b int) {
	p[p.find(b)] = p.find(a)
	o++
}

func (p pSet) print() {
	fmt.Println(p)
}

// 预处理0->nums的质因数
func genGcds(nums int) [][]int {
	res := make([][]int, nums+1)
	for i := range res {
		res[i] = make([]int, 0)
	}

	for i := 2; i <= nums/2; i++ {
		for j := i + i; j <= nums; j += i {
			res[j] = append(res[j], i)
		}
	}
	return res
}

func GcdSort(nums []int) bool {
	defer func() { fmt.Println(o) }()

	maxNum := 0
	nums2 := make([]int, 0)
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		nums2 = append(nums2, num)
	}

	// 初始化质因数
	gs := genGcds(maxNum)

	// 求并查集
	pSet := initPSet(maxNum)
	for _, num := range nums {
		for _, g := range gs[num] {
			pSet.merge(num, g)
		}
	}

	fmt.Println(len(nums))

	// 对比
	sort.Ints(nums2)
	for i, num := range nums2 {
		if num == nums[i] {
			o++
			continue
		}
		if pSet.find(num) != pSet.find(nums[i]) {
			o++
			return false
		}
	}

	return true
}
