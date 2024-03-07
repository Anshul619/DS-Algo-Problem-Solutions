package main

import "testing"

func TestMinWindowSubstring(t *testing.T) {
	tests := []struct {
		s   string
		t   string
		out string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for i, v := range tests {
		if minWindow(v.s, v.t) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
