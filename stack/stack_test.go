package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyntaxCheck(t *testing.T) {
	a := assert.New(t)
	t.Parallel()

	var tests = []struct {
		code     string
		expected int
	}{
		{},
		{
			code:     "()",
			expected: 0,
		},
		{
			code:     "(",
			expected: 1,
		},
		{
			code:     ")",
			expected: 2,
		},
		{
			code:     "(]",
			expected: 2,
		},
		{
			code:     "(({]}))",
			expected: 2,
		},
		{
			code:     "var x = a(b(c[d))",
			expected: 2,
		},
	}

	for _, tt := range tests {
		a.Equal(tt.expected, SyntaxCheckSimple(tt.code), tt.code)
	}
}
