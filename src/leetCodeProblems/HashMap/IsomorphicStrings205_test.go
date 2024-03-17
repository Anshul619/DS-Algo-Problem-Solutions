package main

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for i, v := range tests {
		if isIsomorphic(v.s, v.t) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
