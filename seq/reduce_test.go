package seq

import (
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	expected := 15
	seq := slices.Values([]int{1, 2, 3, 4})
	reduceFunc := func(accumulator, v int) int {
		return accumulator + v
	}
	initial := 5
	actual := Reduce(seq, reduceFunc, initial)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
