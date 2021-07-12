package main

func tribonacci1137(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		a := 0
		b := 1
		c := 1
		for i := 2; i < n; i++ {
			a = a + b + c
			a, b, c = b, c, a
		}
		return c
	}
}
