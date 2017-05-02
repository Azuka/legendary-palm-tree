package basic

func lcm(x int, y int) int {
	return x * y / gcd(x, y)
}
