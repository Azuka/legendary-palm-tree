package sort

func InsertionSort(values []int) []int {
	if len(values) < 2 {
		return values
	}

	// fmt.Printf("Insertion sort before: %+v\n", values)

	for i := 0; i < len(values)-1; i++ {
		// First we create a gap (virtual) at i starting from 1 though n-1
		// and store the value of the gap
		gap := i + 1
		val := values[gap]

		// Next we loop through all values to the left of the gap, ie, less than i
		// and shift them right if they're greater than the current value
		for j := i; j >= 0; j-- {
			// nothing to shift
			if values[j] <= val {
				break
			}

			// shift gap left, ie swap values and update the gap location
			// and comparison value
			values[gap], values[j] = values[j], values[gap]
			gap = j
			val = values[gap]
		}
	}

	// fmt.Printf("Insertion sort after: %+v\n", values)

	return values
}
