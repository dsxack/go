package collect

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMapMap(t *testing.T) {
	res := MapMap(
		map[int]int{1: 1, 2: 2},
		func(v int) string { return strconv.Itoa(v * v) },
	).(map[int]string)

	assert.Equal(t, map[int]string{1: "1", 2: "4"}, res)
}
