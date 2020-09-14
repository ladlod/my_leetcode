package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println(restoreIpAddresses(str))
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
