package main

import (
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		s   string
		out int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abcd", 4},
		{"", 0},
		{"abba", 2},
		{"dvdf", 3},
	}

	for i, v := range tests {
		if lengthOfLongestSubstring(v.s) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
