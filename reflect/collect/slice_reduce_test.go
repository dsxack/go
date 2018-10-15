package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceReduce(t *testing.T) {
	res := SliceReduce(
		[]int{1, 2, 3},
		0,
		func(accum, index, value int) int { return accum + value },
	).(int)

	assert.Equal(t, 6, res)
}
