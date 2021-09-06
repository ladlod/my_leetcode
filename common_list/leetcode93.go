package common_list

import (
	"fmt"
	"strconv"
)

// 给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。

// 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/restore-ip-addresses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 回溯
func RestoreIpAddressesV2(s string) []string {
	return restoreIpAddressesV2Inner(s, 4)
}

func restoreIpAddressesV2Inner(s string, n int) []string {
	if len(s) == 0 && n > 0 {
		return []string{}
	} else if n == 1 {
		if s[0] == '0' && len(s) > 1 {
			return []string{}
		}
		v, _ := strconv.Atoi(s)
		if v <= 255 {
			return []string{s}
		} else {
			return []string{}
		}
	}

	res := make([]string, 0)
	if s[0] == '0' {
		vs := restoreIpAddressesV2Inner(s[1:], n-1)
		for _, v := range vs {
			res = append(res, "0."+v)
		}
		return res
	} else {
		vs := restoreIpAddressesV2Inner(s[1:], n-1)
		for _, v := range vs {
			res = append(res, s[:1]+"."+v)
		}
		if len(s) >= 2 {
			vs = restoreIpAddressesV2Inner(s[2:], n-1)
			for _, v := range vs {
				res = append(res, s[:2]+"."+v)
			}
		}
		if len(s) >= 3 {
			v, _ := strconv.Atoi(s[:3])
			if v <= 255 {
				vs = restoreIpAddressesV2Inner(s[3:], n-1)
				for _, v := range vs {
					res = append(res, s[:3]+"."+v)
				}
			}
		}
		return res
	}
}

func restoreIpAddresses(s string) []string {
	rtn := make([]string, 0)
	for i := 1; i <= 3; i++ {
		if i >= len(s) {
			continue
		}
		substr := s[0:i]
		if substr[:1] == "0" && len(substr) > 1 {
			continue
		}
		if num, _ := strconv.Atoi(substr); num > 255 {
			continue
		}
		str1 := substr
		fmt.Printf("str1:%s\n", substr)
		for j := 1; j <= 3; j++ {
			if i+j >= len(s) {
				continue
			}
			substr := s[i : i+j]
			if substr[:1] == "0" && len(substr) > 1 {
				continue
			}
			if num, _ := strconv.Atoi(substr); num > 255 {
				continue
			}
			str2 := str1 + "." + substr
			fmt.Printf("str2:%s\n", substr)
			for k := 1; k <= 3; k++ {
				if i+j+k >= len(s) {
					continue
				}
				fmt.Println(i, j, k)
				substr := s[i+j : i+j+k]
				if substr[:1] == "0" && len(substr) > 1 {
					continue
				}
				if num, _ := strconv.Atoi(substr); num > 255 {
					continue
				}
				str3 := str2 + "." + substr
				fmt.Printf("str3:%s\n", substr)
				endstr := s[i+j+k:]
				if endstr[:1] == "0" && len(endstr) > 1 {
					continue
				}
				if num, _ := strconv.Atoi(endstr); num > 255 {
					continue
				}
				str := str3 + "." + endstr
				fmt.Printf("str4:%s\n", endstr)
				rtn = append(rtn, str)
			}
		}
	}

	return rtn
}
