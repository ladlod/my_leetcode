package main

import (
	"fmt"

	"github.com/ladlod/leetcode/week_contest"
)

func main() {
	fmt.Println(week_contest.NumberOfUniqueGoodSubsequences("101"))
}

// 0 1
// - -
// [] ["1"]
// ["0", "01"] ["1"]
// ["0", "01"], ["10", "101", "11", "1"]
