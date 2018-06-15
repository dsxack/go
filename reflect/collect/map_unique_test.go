package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapUnique(t *testing.T) {
	res := MapUnique(map[string]int{"one": 1, "anotherOne": 1, "two": 2}).([]int)

	assert.ElementsMatch(t, res, []int{1, 2})
}
