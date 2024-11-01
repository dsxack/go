package seq2

import "iter"

// Take returns a new sequence that contains the first num elements of the original sequence.
// The new sequence will contain only the first num elements.
func Take[K, V any](seq iter.Seq2[K, V], num int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := num
		for k, v := range seq {
			if i <= 0 {
				break
			}
			ok := yield(k, v)
			if !ok {
				break
			}
			i--
		}
	}
}
