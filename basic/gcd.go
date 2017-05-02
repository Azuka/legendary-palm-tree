package basic

// This calculates the greatest common divisor using the Euclidean algorithm
func gcd(n1 int, n2 int) int {
	var min, max int

	if n1 >= n2 {
		min = n2
		max = n1
	} else {
		min = n2
		max = n1
	}

	for min > 0 {
		temp := min
		min = max % min
		max = temp
	}

	return max
}
