package main

/*
- Leetcode - https://leetcode.com/problems/word-pattern/description
*/
import "testing"

func TestWordPattern(t *testing.T) {
	tests := []struct {
		pattern  string
		s        string
		expected bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
		{"aaa", "aa aa aa aa", false},
		{"he", "unit", false},
	}

	for i, v := range tests {
		if wordPattern(v.pattern, v.s) != v.expected {
			t.Errorf("fail %v", i)
		}
	}
}
