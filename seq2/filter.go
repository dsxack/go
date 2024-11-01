package seq2

import "iter"

// Filter returns a new sequence that contains only the elements of the original sequence that satisfy the filter function.
// The filter function should return true for elements that should be included in the new sequence.
func Filter[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !f(k, v) {
				continue
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
