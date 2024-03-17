package main

import "testing"

func TestIsAnagram1(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for i, v := range tests {
		if isAnagram(v.s, v.t) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
