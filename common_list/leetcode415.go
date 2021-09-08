package common_list

// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回。

// 你不能使用任何內建的用于处理大整数的库（比如 BigInteger）， 也不能直接将输入的字符串转换为整数形式。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-strings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func AddStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	res := ""
	add := 0
	for i >= 0 || j >= 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}
		if j >= 0 {
			n2 = int(num2[j] - '0')
		}
		res = string(byte((n1+n2+add)%10)+'0') + res
		add = (n1 + n2 + add) / 10
		i--
		j--
	}
	if add != 0 {
		res = string(byte(add)+'0') + res
	}
	return res
}
