package main

import (
	"fmt"

	"github.com/ladlod/leetcode/daily"
)

func main() {
	v := []int{7, 4, 1, 2, 3, 5}
	daily.QuickSort(v)
	fmt.Println(v)
}
