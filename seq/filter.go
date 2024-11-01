package seq

import "iter"

// Filter returns a new sequence that contains only the elements of the original sequence that satisfy the filter function.
// The filter function should return true for elements that should be included in the new sequence.
func Filter[V any](seq iter.Seq[V], filterFunc func(i V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if filterFunc(v) {
				ok := yield(v)
				if !ok {
					break
				}
			}
		}
	}
}
