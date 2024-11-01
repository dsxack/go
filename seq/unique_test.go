package seq

import (
	"slices"
	"testing"
)

func TestUnique(t *testing.T) {
	expected := []int{1, 2, 3}
	seq := slices.Values([]int{1, 2, 3, 2, 1, 3})

	unique := Unique(seq)
	got := slices.Collect(unique)

	if !slices.Equal(expected, got) {
		t.Errorf("Unique() = %v; want %v", got, expected)
	}
}
