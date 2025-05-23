package main

import "testing"

/*
- LeetCode - https://leetcode.com/problems/majority-element/
*/
func TestMinSwap(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"xx", "yy", 1},
		{"xy", "yx", 2},
		{"xx", "yx", -1},
		{"xx", "xy", -1},
		{"x", "y", -1},
		{"xxyyxyxyxx", "xyyxyxxxyx", 4},
	}

	for i, v := range tests {
		if minimumSwap(v.s1, v.s2) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
