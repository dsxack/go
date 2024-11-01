package seq

import "iter"

// Reduce reduces the sequence to a single value by applying the reduce function to each element.
// The reduce function should take the accumulator and the current element and return the new accumulator.
func Reduce[V any, R any](seq iter.Seq[V], reduceFunc func(R, V) R, initial R) R {
	accumulator := initial
	for v := range seq {
		accumulator = reduceFunc(accumulator, v)
	}
	return accumulator
}
