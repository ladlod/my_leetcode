package common_list

// 给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。

// 进阶：你可以在 O(n) 的时间解决这个问题吗？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 暴力做法：枚举全部组合，获得最大值，时间复杂度O(n^2)

// map做法：
// 异或运算的性质：对于a,b,c，满足a^b=c，那么abc也应当满足a^c=b
// 那么，对于一组n位的数字来说，我们可以遍历0~2^n-1，令该数值为c，判断是否有两个数字进行XOR运算可以得到c，即将所有数字放入一个map中，
// 遍历所有数字，与c进行XOR运算，判断结果是否在map中。对于一个32位数字来说，该做法时间复杂度为O(n*(2^32-1))
// 对此算法做简化：可按位置进行结果计算，先对判断首位需为1或0，后续则只需遍历该位为1的数字即可，以此后推可得：
// 对所有数字的第一位与1取异或，判断是否第一位是否可以为1，若可以，对所有数字的前两位与11取异或，否则与01取异或，判断第二位是否可以为1，以此类推。
// 时间复杂度为O(n*log(2^32-1))

// eg: 3, 10, 5, 25, 2, 8
// 00011, 01010, 00101, 11001, 00010, 01000

func FindMaximumXOR(nums []int) int {
	res := 0
	for i := 31; i >= 0; i-- {
		res <<= 1
		tMap := make(map[int]bool)
		for _, num := range nums {
			tMap[num>>i] = true
		}
		for k := range tMap {
			if _, ok := tMap[k^(res+1)]; ok {
				res += 1
				break
			}
		}
		// fmt.Println(res)
	}

	return res
}
