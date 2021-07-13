package common_list

import (
	"strconv"
	"strings"
)

func numDecodings(s string) int {
	l := len(s)
	if l == 0 || s[:1] == "0" || strings.Contains(s, "00") {
		return 0
	}

	rtn := make([]int, l)
	if s[l-1:] == "0" {
		rtn[l-1] = 0
	} else {
		rtn[l-1] = 1
	}

	if l >= 2 {
		if s[l-2:l-1] == "0" {
			rtn[l-2] = 0
		} else {
			num, _ := strconv.ParseInt(s[l-2:], 10, 64)
			if num > 0 && num < 27 {
				rtn[l-2] = rtn[l-1] + 1
			} else {
				rtn[l-2] = rtn[l-1]
			}
		}
	}

	for i := l - 3; i >= 0; i-- {
		num, _ := strconv.ParseInt(s[i:i+2], 10, 64)
		if num%10 == 0 {
			if num > 0 && num < 27 {
				rtn[i] = rtn[i+2]
			} else {
				return 0
			}
		} else {
			if num > 0 && num < 27 {
				rtn[i] = rtn[i+1] + rtn[i+2]
			} else {
				rtn[i] = rtn[i+1]
			}
		}
	}

	return rtn[0]
}
