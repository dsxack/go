package seq2

import (
	"maps"
	"testing"
)

func TestFilter(t *testing.T) {
	expected := map[int]int{2: 4, 4: 8}
	seq := func(yield func(int, int) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i, i*2) {
				return
			}
		}
	}

	filtered := Filter(seq, func(k, v int) bool {
		return k%2 == 0
	})
	got := maps.Collect(filtered)

	if !maps.Equal(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
