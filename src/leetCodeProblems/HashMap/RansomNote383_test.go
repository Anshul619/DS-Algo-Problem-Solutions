package main

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for i, v := range tests {
		if canConstruct(v.ransomNote, v.magazine) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
