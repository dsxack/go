package collect

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapMapWithKeys(t *testing.T) {
	res := MapMapWithKeys(
		map[int]int{1: 1, 2: 2},
		func(k, v int) (string, string) {
			return strconv.Itoa(k), strconv.Itoa(v * v)
		},
	).(map[string]string)

	assert.Equal(t, map[string]string{"1": "1", "2": "4"}, res)
}
