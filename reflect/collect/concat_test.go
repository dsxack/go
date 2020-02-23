package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat_SameTypes(t *testing.T) {
	values := Concat([]int{1, 2}, []int{3, 4}, []int{5, 6}).([]int)

	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, values)
}

func TestConcat_DifferentTypes(t *testing.T) {
	values := Concat([]int{1, 2}, []int{3, 4}, []float64{5, 6}).([]interface{})

	assert.EqualValues(t, []interface{}{1, 2, 3, 4, float64(5), float64(6)}, values)
}

func TestConcat_OneArg(t *testing.T) {
	values := Concat([][]int{
		{1, 2}, {3, 4}, {5, 6},
	}).([]int)

	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, values)
}
