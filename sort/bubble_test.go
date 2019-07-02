package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	a := assert.New(t)

	for _, tt := range tests() {
		a.Equal(tt.expected, BubbleSort(tt.original), tt.name)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	vals := make([]int, b.N)

	for i := 0; i < len(vals); i++ {
		vals[i] = len(vals) - i
	}

	BubbleSort(vals)
}
