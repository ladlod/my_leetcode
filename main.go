package main

import (
	"fmt"

	"github.com/ladlod/leetcode/week_contest"
)

func main() {
	test1, test2 := []int{}, []int{}
	for i := 10000; i > 0; i-- {
		test1 = append(test1, i)
	}
	for i := 1; i <= 10000; i++ {
		test2 = append(test2, i)
	}
	fmt.Println(week_contest.GcdSort(test1))
	fmt.Println(week_contest.GcdSort(test2))
}
