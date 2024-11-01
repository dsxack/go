package seq2

import (
	"maps"
	"slices"
	"testing"
)

func TestCollectSlice(t *testing.T) {
	expected := []int{1, 2, 3}
	seq := func(yield func(string, int) bool) {
		yield("one", 1)
		yield("two", 2)
		yield("three", 3)
	}

	got := CollectSlice(seq)

	if !slices.Equal(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}

func TestCollectMap(t *testing.T) {
	seq := func(yield func(string, int) bool) {
		yield("one", 1)
		yield("two", 2)
		yield("three", 3)
	}
	expected := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	got := CollectMap(seq)

	if !maps.Equal(got, expected) {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
