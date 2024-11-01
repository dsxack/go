package seq

import (
	"maps"
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	expected := []int{2, 3, 4, 5, 6}

	mapped := Map(seq, func(v int) int {
		return v + 1
	})

	got := slices.Collect(mapped)

	if !slices.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestMap2(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	expected := map[int]int{1: 2, 2: 4, 3: 6, 4: 8, 5: 10}

	mapped := Map2(seq, func(v int) (int, int) {
		return v, v * 2
	})

	got := maps.Collect(mapped)

	if !maps.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
