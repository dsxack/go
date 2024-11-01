package seq

import (
	"slices"
	"testing"
)

func TestTake(t *testing.T) {
	expected := []int{1, 2, 3}
	seq := slices.Values([]int{1, 2, 3, 4, 5})

	taken := Take(seq, 3)
	got := slices.Collect(taken)

	if !slices.Equal(expected, got) {
		t.Errorf("Take() = %v; want %v", got, expected)
	}
}
