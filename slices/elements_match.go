package slices

import "maps"

// ElementsMatch returns true if the two slices contain the same elements, regardless of order.
// The elements must be comparable.
func ElementsMatch[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1set := make(map[E]int, len(s1))
	for _, v := range s1 {
		s1set[v]++
	}
	s2set := make(map[E]int, len(s2))
	for _, v := range s2 {
		s2set[v]++
	}
	return maps.Equal(s1set, s2set)
}
