package collect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapKeys(t *testing.T) {
	res := MapKeys(map[string]int{
		"one": 1,
		"two": 2,
	}).([]string)

	assert.ElementsMatch(t, res, []string{"one", "two"})
}
