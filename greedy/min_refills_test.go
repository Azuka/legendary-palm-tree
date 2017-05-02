package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var minRefillsTests = []struct {
	input    []int
	distance int
	expected int
}{
	{
		[]int{0, 200, 375, 550, 750, 950},
		400,
		2,
	},
	{
		[]int{0, 200, 375, 550, 750, 950},
		300,
		4,
	},
	{
		[]int{0, 200, 375, 550, 750, 950},
		500,
		2,
	},
	{
		[]int{0, 200, 375, 550, 750, 950},
		600,
		1,
	},
	{
		[]int{0, 200, 375, 550, 750, 950},
		950,
		0,
	},
}

func TestMinRefills(t *testing.T) {
	assert := assert.New(t)
	for _, tt := range minRefillsTests {
		assert.Equal(tt.expected, minRefills(tt.input, tt.distance))
	}
}
