package sort

func SelectionSort(values []int) []int {

	if len(values) <= 1 {
		return values
	}

	for i := 0; i < len(values); i++ {
		min := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[min] {
				min = j
			}
		}

		values[min], values[i] = values[i], values[min]
		// fmt.Printf("Selection sort pass %d val = %+v\n", i, values)
	}

	return values
}
