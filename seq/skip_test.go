package seq

import (
	"slices"
	"testing"
)

func TestSkip(t *testing.T) {
	expected := []int{4, 5}
	seq := slices.Values([]int{1, 2, 3, 4, 5})

	skipped := Skip(seq, 3)
	got := slices.Collect(skipped)

	if !slices.Equal(expected, got) {
		t.Errorf("Skip() = %v; want %v", got, expected)
	}
}
