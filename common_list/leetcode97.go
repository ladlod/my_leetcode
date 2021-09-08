package common_list

// 给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。

// 两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：

// s = s1 + s2 + ... + sn
// t = t1 + t2 + ... + tm
// |n - m| <= 1
// 交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
// 提示：a + b 意味着字符串 a 和 b 连接。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/interleaving-string
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路
// 动态规划
// 设dp[i][j]表示 s1[i] + s2[j] == s3[i+j]
// 则 dp[i][j] = (dp[i-1][j] && s1[i] == s3[i+j]) || (dp[i][j-1] && s2[j] == s3[i+j]])
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	} else {
		res := make([][]bool, len(s1)+1)
		for i := 0; i <= len(s1); i++ {
			res[i] = make([]bool, len(s2)+1)
			res[i][0] = s1[:i] == s3[:i]
		}
		for j := 0; j <= len(s2); j++ {
			res[0][j] = s2[:j] == s3[:j]
		}

		for i := 0; i < len(s1); i++ {
			for j := 0; j < len(s2); j++ {
				if res[i][j+1] {
					if s1[i] == s3[i+j+1] {
						res[i+1][j+1] = true
					}
				} else if res[i+1][j] {
					if s2[j] == s3[i+j+1] {
						res[i+1][j+1] = true
					}
				} else {
					res[i+1][j+1] = false
				}
			}
		}
		return res[len(s1)][len(s2)]
	}
}
