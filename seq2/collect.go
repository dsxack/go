package seq2

import (
	"cmp"
	"iter"
	"maps"
)

// CollectSlice collects all elements of the sequence into a slice.
func CollectSlice[K, V any](seq iter.Seq2[K, V]) []V {
	var result []V
	for _, v := range seq {
		result = append(result, v)
	}
	return result
}

// CollectMap collects all elements of the sequence into a map.
func CollectMap[K cmp.Ordered, V any](seq iter.Seq2[K, V]) map[K]V {
	return maps.Collect(seq)
}
