package seq2

import (
	"maps"
	"testing"
)

func TestTake(t *testing.T) {
	expected := map[int]int{
		1: 2,
		2: 4,
	}
	seq := func(yield func(int, int) bool) {
		for i := 1; i <= 4; i++ {
			if !yield(i, i*2) {
				return
			}
		}
	}

	taken := Take(seq, 2)
	got := maps.Collect(taken)

	if !maps.Equal(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
