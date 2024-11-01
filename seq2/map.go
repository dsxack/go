package seq2

import "iter"

// Map returns a new sequence that contains the results of applying the mapper function to each element of the original sequence.
// The mapper function should return the new value for each element.
func Map[InK, InV, OutV any](
	seq iter.Seq2[InK, InV],
	mapper func(InK, InV) OutV,
) iter.Seq[OutV] {
	return func(yield func(OutV) bool) {
		for k, v := range seq {
			ok := yield(mapper(k, v))
			if !ok {
				break
			}
		}
	}
}

// Map2 returns a new sequence that contains the result of applying the mapper function to each element of the original sequence.
// The mapper function should return a key-value pair for each element.
func Map2[InK, InV, OutK, OutV any](
	seq iter.Seq2[InK, InV],
	mapper func(InK, InV) (OutK, OutV),
) iter.Seq2[OutK, OutV] {
	return func(yield func(OutK, OutV) bool) {
		for k, v := range seq {
			ok := yield(mapper(k, v))
			if !ok {
				break
			}
		}
	}
}
