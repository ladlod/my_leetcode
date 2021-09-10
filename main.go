package main

import (
	"fmt"
)

func main() {
	test := []int{}
	for i := 1; i < 5; i++ {
		test = append(test, i)
		fmt.Println(len(test), cap(test))
	}

	test1 := []int{1}
	test2 := []int{1, 2}
	test3 := []int{1, 2, 3}
	test4 := []int{1, 2, 3, 4}
	fmt.Println(len(test1), cap(test1))
	fmt.Println(len(test2), cap(test2))
	fmt.Println(len(test3), cap(test3))
	fmt.Println(len(test4), cap(test4))
}
