// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

// P   A   H   N
// A P L S I I G
// Y   I   R

// P     I     N
// A   L S   I G
// Y A   H R
// P     I

// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

// 请你实现这个将字符串进行指定行数变换的函数：

// string convert(string s, int numRows);

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/zigzag-conversion
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 第i行的数字由列头-行号和列头+行号组成。循环遍历即可，首行和尾行只加一次
// 需要注意边界问题
package common_list

import (
	"bytes"
	"fmt"
)

func Convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	var res bytes.Buffer
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s)+i; j += (2*numRows - 2) {
			if i == 0 {
				res.WriteByte(s[j-i])
			} else if i > 0 && i < numRows-1 {
				if j-i >= 0 {
					res.WriteByte(s[j-i])
				}
				if j+i < len(s) {
					res.WriteByte(s[j+i])
				}
			} else if i == numRows-1 {
				if j+i < len(s) {
					res.WriteByte(s[j+i])
				}
			}
		}
		fmt.Println(res.String())
	}

	return res.String()
}
