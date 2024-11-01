package seq

import (
	"slices"
	"testing"
)

func TestRange(t *testing.T) {
	seq := Range[int](0, 10)
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := slices.Collect(seq)

	if !slices.Equal(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
