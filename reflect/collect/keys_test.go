package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	res := Keys(map[string]int{
		"one": 1,
		"two": 2,
	}).([]string)

	assert.ElementsMatch(t, res, []string{"one", "two"})
}
