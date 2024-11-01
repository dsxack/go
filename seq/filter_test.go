package seq

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	expected := []int{2, 4, 6, 8}
	seq := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8})

	filtered := Filter(seq, func(i int) bool {
		return i%2 == 0
	})
	got := slices.Collect(filtered)

	if !slices.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
