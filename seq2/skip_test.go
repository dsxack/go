package seq2

import (
	"maps"
	"testing"
)

func TestSkip(t *testing.T) {
	expect := map[int]int{
		3: 6,
		4: 8,
		5: 10,
	}
	seq := func(yield func(int, int) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, i*2) {
				return
			}
		}
	}

	skipped := Skip(seq, 2)
	got := maps.Collect(skipped)

	if !maps.Equal(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}
}
