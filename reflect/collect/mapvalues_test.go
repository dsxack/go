package collect

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapValues_Map(t *testing.T) {
	res := MapValues(
		map[int]int{1: 1, 2: 2},
		func(_, v int) string { return strconv.Itoa(v * v) },
	).(map[int]string)

	assert.Equal(t, map[int]string{1: "1", 2: "4"}, res)
}

func TestMapValues_Slice(t *testing.T) {
	res := MapValues(
		[]int{1, 2, 3},
		func(_, v int) string {
			return strconv.Itoa(v * v)
		},
	).([]string)

	assert.ElementsMatch(t, []string{"1", "4", "9"}, res)
}
