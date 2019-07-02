package dnc

func multiplyPolynomialsNaive(a []int, b []int, n int) []int {
	l := (n * 2) - 1
	res := make([]int, l)

	for i := range a {
		for j := range b {
			res[i+j] += a[i] * b[j]
		}
	}

	return res
}

func mult2(a []int, b []int, n int, subA int, subB int) []int {

	r := make([]int, 2*n-1)

	if n == 1 {
		r[0] = a[subA] * b[subB]
		return r
	}

	half := n / 2

	first := mult2(a, b, half, subA, subB)
	second := mult2(a, b, half, subA+half, subB+half)

	for i := 0; i <= n-2; i++ {
		r[i] = first[i]
		r[i+n] = second[i]
	}

	d0e1 := mult2(a, b, half, subA, subB+half)
	d1e0 := mult2(a, b, half, subA+half, subB)

	sum := make([]int, len(d0e1))

	for i := range d0e1 {
		sum[i] = d0e1[i] + d1e0[i]
	}

	for i := 0; i <= n-2; i++ {
		r[half+i] += sum[i]
	}

	return r
}

func multiplyPolynomialsSlightlyBetter(a []int, b []int, n int) []int {
	return mult2(a, b, n, 0, 0)
}
