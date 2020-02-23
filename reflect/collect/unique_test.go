package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique_Map(t *testing.T) {
	res := Unique(map[string]int{"one": 1, "anotherOne": 1, "two": 2}).([]int)

	assert.ElementsMatch(t, res, []int{1, 2})
}

func TestUnique_Slice(t *testing.T) {
	res := Unique([]int{1, 1, 2}).([]int)

	assert.ElementsMatch(t, res, []int{1, 2})
}
