package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fibonacciTests = []struct {
	n        int
	expected int
}{
	{
		0,
		0,
	},
	{
		1,
		1,
	},
	{
		2,
		1,
	},
	{
		3,
		2,
	},
	{
		4,
		3,
	},
	{
		63,
		6557470319842,
	},
}

func TestFibonacci(t *testing.T) {
	assert := assert.New(t)

	for _, tt := range fibonacciTests {
		assert.Equal(tt.expected, fibonacci(tt.n), fmt.Sprintf("fibonacci(%d) expected %d", tt.n, tt.expected))
	}
}
