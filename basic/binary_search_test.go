package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var binarySearchTests = []struct {
	haystack      []int
	needle        int
	foundPosition int
	expectedError error
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		1,
		0,
		nil,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		2,
		1,
		nil,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		2,
		nil,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		4,
		3,
		nil,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		5,
		4,
		nil,
	},
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		8,
		-1,
		NotFound,
	},
}

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)
	for _, tt := range binarySearchTests {
		pos, err := binarySearch(tt.haystack, tt.needle)
		assert.Equal(tt.foundPosition, pos)
		assert.Equal(tt.expectedError, err)
	}
}
