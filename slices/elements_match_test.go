package slices

import "testing"

func TestElementsMatch(t *testing.T) {
	tests := []struct {
		name     string
		s1       []int
		s2       []int
		expected bool
	}{
		{
			name:     "s1 and s2 are equal",
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "s1 and s2 are not equal",
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 4},
			expected: false,
		},
		{
			name:     "s1 and s2 have different lengths",
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "s1 and s2 are empty",
			s1:       nil,
			s2:       nil,
			expected: true,
		},
		{
			name:     "s1 and s2 have the same elements but in different order",
			s1:       []int{1, 2, 2},
			s2:       []int{2, 2, 1},
			expected: true,
		},
		{
			name:     "s1 and s2 have the same elements but different counts",
			s1:       []int{1, 2, 2},
			s2:       []int{1, 2, 1},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ElementsMatch(tt.s1, tt.s2); got != tt.expected {
				t.Errorf("ElementsMatch() = %v, want %v", got, tt.expected)
			}
		})
	}
}
