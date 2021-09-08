package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	fmt.Println(common_list.ExistStr([][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'},
	}, "abccedasf"))
}
