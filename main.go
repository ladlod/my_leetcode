package main

import (
	"fmt"

	"github.com/ladlod/leetcode/dp"
)

func main() {
	fmt.Println(dp.WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
