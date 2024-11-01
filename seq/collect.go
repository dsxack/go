package seq

import (
	"iter"
	"slices"
)

// CollectSlice collects all elements from the sequence into a slice.
// It is just a wrapper around slices.Collect.
func CollectSlice[V any](seq iter.Seq[V]) []V {
	return slices.Collect(seq)
}
