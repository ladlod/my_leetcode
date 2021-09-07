package main

import (
	"fmt"

	"github.com/ladlod/leetcode/common_list"
)

func main() {
	fmt.Println(common_list.NumIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		// {"1", "1", "0", "0", "0"},
		// {"0", "0", "0", "0", "0"},
	}))
}
