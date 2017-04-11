package basic

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	l := 0
	r := 1

	for i := 1; i < n; i++ {
		temp := r
		r = l + r
		l = temp
	}

	return r
}
