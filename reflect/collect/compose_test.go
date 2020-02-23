package collect

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestCompose(t *testing.T) {
	f := Compose(
		strconv.Itoa,
		math.Abs,
		FnReduce(0, func(acc, _, value int) int { return acc - value }),
		Keys,
		FnMapWithKeys(func(k, v int) (int, int) { return v, k }),
		FnMapValues(func(k, v int) int { return v * 2 }),
		Unique,
		func(values ...int) []int { return values },
	)

	result := f(1, 2, 3, 3, 4).(string)

	assert.Equal(t, "20", result)
}

func TestCompose2(t *testing.T) {
	classyGreeting := func(firstName, lastName string) string {
		return "The name's " + lastName + ", " + firstName + " " + lastName
	}

	yellGreeting := Compose(strings.ToUpper, classyGreeting)

	assert.Equal(t, "THE NAME'S BOND, JAMES BOND", yellGreeting("James", "Bond"))
}

func TestCompose3(t *testing.T) {
	multiply := func(a int) func(int) int {
		return func(b int) int {
			return a * b
		}
	}

	add := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	f := Compose(math.Abs, add(1), multiply(2))

	assert.Equal(t, float64(7), f(-4))
}
