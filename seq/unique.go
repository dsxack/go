package seq

import "iter"

// Unique returns a new sequence that contains only the unique elements of the original sequence.
// The new sequence will contain only the first occurrence of each element.
func Unique[E comparable](s iter.Seq[E]) iter.Seq[E] {
	seen := make(map[E]struct{})
	return Filter(s, func(e E) bool {
		if _, ok := seen[e]; ok {
			return false
		}
		seen[e] = struct{}{}
		return true
	})
}
