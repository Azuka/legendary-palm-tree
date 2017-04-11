package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var maxPairwiseTests = []struct {
	array         []int
	expectedValue int
}{
	{
		[]int{1, 2, 3, 4, 5},
		20,
	},
	{
		[]int{1, 2, 30, 4, 5},
		150,
	},
	{
		[]int{1, 2, 30, 30, 5},
		900,
	},
	{
		generateSequence(1, 100000),
		9999900000,
	},
}

func generateSequence(min int, num int) []int {
	sequence := make([]int, num)
	for i := 0; i < num; i++ {
		sequence[i] = min + i
	}
	return sequence
}

func TestMaxPairwiseProduct(t *testing.T) {
	assert := assert.New(t)

	for _, tt := range maxPairwiseTests {
		assert.Equal(tt.expectedValue, maxPairwiseProduct(tt.array))
	}
}
