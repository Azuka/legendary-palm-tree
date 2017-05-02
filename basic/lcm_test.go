package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lcmTests = []struct {
	x        int
	y        int
	expected int
}{
	{
		4,
		6,
		12,
	},
	{
		10,
		15,
		30,
	},
}

func TestLcm(t *testing.T) {
	assert := assert.New(t)

	for _, tt := range lcmTests {
		assert.Equal(tt.expected, lcm(tt.x, tt.y))
		assert.Equal(tt.expected, lcm(tt.y, tt.x))
	}
}
