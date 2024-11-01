package seq

import (
	"github.com/dsxack/go/v3/constraints"
	"iter"
)

// Range returns a new sequence that contains all integers from start to end (exclusive).
// The start value is included in the sequence, but the end value is not.
func Range[V constraints.Integer](start, end V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := start; i < end; i++ {
			ok := yield(i)
			if !ok {
				break
			}
		}
	}
}
