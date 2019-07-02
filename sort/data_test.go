package sort

type test struct {
	original []int
	expected []int
	name     string
}

func tests() []test {
	return []test{

		{
			original: []int{4, 2, 1, 7, 3},
			expected: []int{1, 2, 3, 4, 7},
			name:     "random",
		},
		{
			original: []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
			name:     "sorted asc",
		},
		{
			original: []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
			name:     "sorted desc",
		},
	}
}
