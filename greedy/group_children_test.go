package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var groupChildrenTests = []struct {
	ages               []float32
	span               float32
	expectedGroupCount int
}{
	{
		[]float32{1.0, 1.5, 3.0, 3.5, 4.0, 7.0},
		1,
		3,
	},
	{
		[]float32{1.0, 1.5, 3.0, 3.5, 4.0, 7.0},
		1.5,
		3,
	},
	{
		[]float32{1.0, 3.0, 7.0},
		1,
		3,
	},
}

func TestGroupChildren(t *testing.T) {
	assert := assert.New(t)

	for _, tt := range groupChildrenTests {
		assert.Equal(tt.expectedGroupCount, minChildGroups(tt.ages, tt.span))
	}
}
