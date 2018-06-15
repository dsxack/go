package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceUnique(t *testing.T) {
	res := SliceUnique([]int{1, 1, 2}).([]int)

	assert.ElementsMatch(t, res, []int{1, 2})
}
