package dnc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	a        []int
	b        []int
	n        int
	expected []int
}{
	{
		a:        []int{3, 4},
		b:        []int{1, 2},
		n:        2,
		expected: []int{3, 10, 8},
	},
	{
		a:        []int{1, 1},
		b:        []int{1, -1},
		n:        2,
		expected: []int{1, 0, -1},
	},
	{
		a:        []int{4, 3, 2, 1},
		b:        []int{1, 2, 3, 4},
		n:        4,
		expected: []int{4, 11, 20, 30, 20, 11, 4},
	},
}

func TestMultiplyPolynomialsNaive(t *testing.T) {
	a := assert.New(t)

	for _, tt := range testCases {
		result := multiplyPolynomialsNaive(tt.a, tt.b, tt.n)
		a.Lenf(result, (2*tt.n)-1, "failed: a:%+v, b:%+v", tt.a, tt.b)
		a.Equalf(tt.expected, result, "failed: a:%+v, b:%+v", tt.a, tt.b)
	}
}

func TestMultiplyPolynomialsSlightlyBetter(t *testing.T) {
	a := assert.New(t)

	for _, tt := range testCases {
		result := multiplyPolynomialsSlightlyBetter(tt.a, tt.b, tt.n)
		a.Lenf(result, (2*tt.n)-1, "len(a) %d != %d failed: a:%+v, b:%+v", tt.a, tt.b)
		a.Equalf(tt.expected, result, "failed: a:%+v, b:%+v", tt.a, tt.b)
	}
}

func BenchmarkMultiplyPolynomialNaive(bt *testing.B) {
	a := make([]int, bt.N)
	b := make([]int, bt.N)
	n := bt.N

	multiplyPolynomialsNaive(a, b, n)
}

func BenchmarkMultiplyPolynomialSlightlyBetter(bt *testing.B) {
	n := int(math.Pow(2, math.Ceil(math.Log2(float64(bt.N)))))

	a := make([]int, n)
	b := make([]int, n)

	multiplyPolynomialsSlightlyBetter(a, b, n)
}
