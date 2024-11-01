package seq

import (
	"slices"
	"testing"
)

func TestCollectSlice(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	seq := slices.Values([]int{1, 2, 3, 4, 5})

	got := CollectSlice(seq)

	if !slices.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
