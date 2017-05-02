package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var gcdTests = []struct {
	x        int
	y        int
	expected int
}{
	{
		0,
		0,
		0,
	},
	{
		0,
		1,
		1,
	},
	{
		4,
		6,
		2,
	},
	{
		4000,
		6000,
		2000,
	},
}

func TestGcd(t *testing.T) {
	assert := assert.New(t)

	for _, tt := range gcdTests {
		assert.Equal(tt.expected, gcd(tt.x, tt.y))
		assert.Equal(tt.expected, gcd(tt.y, tt.x))
	}
}
