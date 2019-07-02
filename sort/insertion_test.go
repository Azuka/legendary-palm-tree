package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {

	a := assert.New(t)

	for _, tt := range tests() {
		a.Equal(tt.expected, InsertionSort(tt.original), tt.name)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	vals := make([]int, b.N)

	for i := 0; i < len(vals); i++ {
		vals[i] = len(vals) - i
	}

	InsertionSort(vals)
}
