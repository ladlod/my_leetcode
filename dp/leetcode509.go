package dp

func fib509(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		a := 0
		b := 1
		for i := 1; i < n; i++ {
			a = a + b
			a, b = b, a
		}
		return b
	}
}
