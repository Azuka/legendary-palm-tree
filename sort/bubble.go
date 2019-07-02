package sort

// BubbleSort checks every item in the slice
// and moves it one right if it is larger than the target to the right
// doing this n*n times ensures all the large items bubble to the top
func BubbleSort(values []int) []int {

	if len(values) <= 1 {
		return values
	}

	// We pass through as many times as we have items in the array
	for x := 0; x < len(values); x++ {
		// For each passthrough, we know we have to check one less item in the list because the larget
		// number bubbles to the top
		for j := 1; j < len(values)-x; j++ {
			i := j - 1
			if values[i] > values[j] {
				values[j], values[i] = values[i], values[j]
			}
		}
		// fmt.Printf("BuubleSort passthrough %d val = %+v\n", x, values)
	}

	return values
}
