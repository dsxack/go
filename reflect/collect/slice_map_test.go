package collect

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSliceMap(t *testing.T) {
	res := SliceMap(
		[]int{1, 2, 3},
		func(v int) string {
			return strconv.Itoa(v * v)
		},
	).([]string)

	assert.ElementsMatch(t, res, []string{"1", "4", "9"})
}
