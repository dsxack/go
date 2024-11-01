package seq2

import (
	"maps"
	"slices"
	"testing"
)

func TestMap(t *testing.T) {
	expected := []uint{0, 6, 12, 18, 24}
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 5; i++ {
			if !yield(i, i*2) {
				return
			}
		}
	}

	mapped := Map(seq, func(_ int, v int) uint { return uint(v * 3) })
	got := slices.Collect(mapped)

	if !slices.Equal(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestMap2(t *testing.T) {
	expected := map[int]uint{
		0: 0,
		1: 6,
		2: 12,
		3: 18,
		4: 24,
	}
	seq := func(yield func(int, int) bool) {
		for i := 0; i < 5; i++ {
			if !yield(i, i*2) {
				return
			}
		}
	}

	mapped := Map2(seq, func(k int, v int) (int, uint) { return k, uint(v * 3) })
	got := maps.Collect(mapped)

	if !maps.Equal(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
