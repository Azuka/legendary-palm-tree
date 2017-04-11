package basic

func maxPairwiseProduct(array []int) int {
	max1 := -1
	max2 := -1

	for i, tt := range array {
		if max1 == -1 || tt > array[max1] {
			max1 = i
		}
	}
	for i, tt := range array {
		if i != max1 && (max2 == -1 || tt > array[max2]) {
			max2 = i
		}
	}

	return array[max1] * array[max2]
}
