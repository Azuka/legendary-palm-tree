package greedy

func minRefills(input []int, maxDistance int) int {
	num := 0
	lastRefill := 0

	for lastRefill < (len(input) - 1) {
		for next := lastRefill; next < (len(input)); next++ {
			if input[next]-input[lastRefill] > maxDistance {
				lastRefill = next - 2
				num++
				break
			}
		}
		lastRefill++
	}

	return num
}
