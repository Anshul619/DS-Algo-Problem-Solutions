package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for i, v := range tests {
		if longestCommonPrefix(v.strs) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
