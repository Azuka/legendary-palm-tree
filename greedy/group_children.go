package greedy

// Given a party of children with ages in sorted order,
// find the minimum number of groups that can be formed of children
// within a particular age limit
func minChildGroups(childAges []float32, limit float32) int {

	count := 1
	min := childAges[0]

	for i := 0; i < len(childAges); i++ {
		if childAges[i]-min > limit {
			count++
			min = childAges[i]
		}
	}

	return count
}
