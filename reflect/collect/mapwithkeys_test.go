package collect

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapWithKeys_Map(t *testing.T) {
	res := MapWithKeys(
		map[int]int{1: 1, 2: 2},
		func(k, v int) (string, string) {
			return strconv.Itoa(k), strconv.Itoa(v * v)
		},
	).(map[string]string)

	assert.Equal(t, map[string]string{"1": "1", "2": "4"}, res)
}

func TestMapWithKeys_Slice(t *testing.T) {
	res := MapWithKeys(
		[]int{1, 2, 3},
		func(k, v int) (int, string) {
			return v, strconv.Itoa(v * v)
		},
	).(map[int]string)

	assert.Equal(t, map[int]string{1: "1", 2: "4", 3: "9"}, res)
}
