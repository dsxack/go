package seq

import "iter"

// Take returns a new sequence that contains the first num elements of the original sequence.
// The new sequence will contain only the first num elements.
func Take[V any](seq iter.Seq[V], num int) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := num
		for v := range seq {
			if i <= 0 {
				break
			}
			ok := yield(v)
			if !ok {
				break
			}
			i--
		}
	}
}
