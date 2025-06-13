package main

import "testing"

func Test_longestSubstringWithAtMostKUniqueCharacters340(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"eceba", 2, 3},
		{"aa", 1, 2},
		{"aabbcc", 2, 4},
		{"abcadcacacaca", 3, 11},
		{"abc", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestSubstringWithAtMostKUniqueCharacters340(tt.s, tt.k); got != tt.expected {
				t.Errorf("longestSubstringWithAtMostKUniqueCharacters340() = %v, want %v", got, tt.expected)
			}
		})
	}
}
