package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapReduce(t *testing.T) {
	res := MapReduce(
		map[int]int{1: 1, 2: 2},
		0,
		func(accum int, key, value int) int { return accum + value },
	).(int)

	assert.Equal(t, 3, res)
}
