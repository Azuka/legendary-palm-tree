package sort

func QuickSort(input []int) []int {
	return quickSortHelper(input, 0, len(input)-1)
}

func quickSortHelper(input []int, left int, right int) []int {

	// base case, only one item, or parition bounds have crossed
	if right <= left {
		return input
	}

	// get array and pivot from partition helper
	input, pivot := quickPartitionHelper(input, left, right)

	// recursively sort the left...
	input = quickSortHelper(input, left, pivot-1)
	// ...then the right
	input = quickSortHelper(input, pivot+1, right)

	return input
}

func quickPartitionHelper(input []int, left int, right int) ([]int, int) {

	// set pivot to right-most element
	pivot := right
	// then set bounds to one less than right
	right--

	for {
		// move left until it's >= pivot
		for input[left] < input[pivot] {
			left++
		}

		// move right until it's <= pivot value, or until it reaches the beginning of the array
		for input[right] > input[pivot] && right > 0 {
			right--
		}

		// stop when left and right have criss-crossed
		if left >= right {
			break
		}

		// swap left and right inputs otherwise
		input[left], input[right] = input[right], input[left]
	}

	// finally, swap pivot and left to get the new pivot
	input[pivot], input[left] = input[left], input[pivot]

	// return parititoned array and pivot position
	return input, left
}

// just a simple test function to verify correctness
func quickPartition(input []int) ([]int, int) {
	left := 0
	right := len(input) - 1

	return quickPartitionHelper(input, left, right)
}
