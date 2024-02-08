package main

import "testing"

func TestMinDeletions(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"aab", 0},
		{"aaabbbcc", 2},
		{"ceabaacb", 2},
		{"bbcebab", 2},
	}

	for i, v := range tests {
		if minDeletions(v.s) != v.expected {
			//t.Fail()
			t.Errorf("failed test %v", i)
		}
	}
}
