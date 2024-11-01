package seq

import "iter"

// Skip returns a new sequence that skips the first num elements of the original sequence.
// The new sequence will contain all elements after the first num elements.
func Skip[V any](seq iter.Seq[V], num uint) iter.Seq[V] {
	return func(yield func(V) bool) {
		var i uint = 0
		for v := range seq {
			if i >= num {
				ok := yield(v)
				if !ok {
					break
				}
			}
			i += 1
		}
	}
}
