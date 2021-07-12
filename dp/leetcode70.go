package main

func climbStairs70(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		a := 1
		b := 2
		for i := 2; i < n; i++ {
			a = a + b
			a, b = b, a
		}
		return b
	}
}
