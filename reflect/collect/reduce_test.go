package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReduce_Slice(t *testing.T) {
	res := Reduce(
		[]int{1, 2, 3},
		0,
		func(accum, index, value int) int { return accum + value },
	).(int)

	assert.Equal(t, 6, res)
}

func TestReduce_Map_Int(t *testing.T) {
	res := Reduce(
		map[int]int{1: 1, 2: 2},
		0,
		func(accum int, key, value int) int { return accum + value },
	).(int)

	assert.Equal(t, 3, res)
}

func TestReduce_Map_Map(t *testing.T) {
	res := Reduce(
		map[int]int{1: 1, 2: 2},
		map[int]int{},
		func(accum map[int]int, key, value int) map[int]int {
			accum[key] = value * 2
			return accum
		},
	).(map[int]int)

	assert.EqualValues(t, map[int]int{1: 2, 2: 4}, res)
}
