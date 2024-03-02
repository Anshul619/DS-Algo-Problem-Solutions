package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for i, v := range tests {
		if isSubsequence(v.s, v.t) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
