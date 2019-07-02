package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {

	a := assert.New(t)

	for _, tt := range tests() {
		a.Equal(tt.expected, QuickSort(tt.original), tt.name)
	}
}

func TestPivot(t *testing.T) {
	a := assert.New(t)
	partition, _ := quickPartition([]int{0, 5, 2, 1, 6, 3})
	a.Equal(
		[]int{0, 1, 2, 3, 6, 5},
		partition,
	)
}

func BenchmarkQuickSort(b *testing.B) {
	vals := make([]int, b.N)

	for i := 0; i < len(vals); i++ {
		vals[i] = len(vals) - i
	}

	QuickSort(vals)
}
