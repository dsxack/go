package seq

import "iter"

// Map returns a new sequence that contains the result of applying the mapper function to each element of the original sequence.
// The mapper function should return the new value for each element.
func Map[In any, Out any](
	seq iter.Seq[In],
	mapper func(In) Out,
) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for v := range seq {
			ok := yield(mapper(v))
			if !ok {
				break
			}
		}
	}
}

// Map2 returns a new sequence that contains the result of applying the mapper function to each element of the original sequence.
// The mapper function should return a key-value pair for each element.
func Map2[InV, OutV, OutK any](
	seq iter.Seq[InV],
	mapper func(InV) (OutK, OutV),
) iter.Seq2[OutK, OutV] {
	return func(yield func(OutK, OutV) bool) {
		for v := range seq {
			ok := yield(mapper(v))
			if !ok {
				break
			}
		}
	}
}
