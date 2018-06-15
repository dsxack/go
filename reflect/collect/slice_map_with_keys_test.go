package collect

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSliceMapWithKeys(t *testing.T) {
	res := SliceMapWithKeys(
		[]int{1, 2, 3},
		func(k, v int) (int, string) {
			return v, strconv.Itoa(v * v)
		},
	).(map[int]string)

	assert.Equal(t, map[int]string{1: "1", 2: "4", 3: "9"}, res)
}
