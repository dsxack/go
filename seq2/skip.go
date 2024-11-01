package seq2

import "iter"

// Skip returns a new sequence that skips the first num elements of the original sequence.
// The new sequence will contain all elements after the first num elements.
func Skip[K, V any](seq iter.Seq2[K, V], num int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range seq {
			if i >= num {
				ok := yield(k, v)
				if !ok {
					break
				}
			}
			i += 1
		}
	}
}
