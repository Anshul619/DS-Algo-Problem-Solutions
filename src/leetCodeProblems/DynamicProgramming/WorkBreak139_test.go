package main

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"", []string{"a"}, true},
		{"a", []string{}, false},
		{"a", []string{"a"}, true},
		{"aaaaaaa", []string{"aaaa", "aaa"}, true},
		{"aaaaaaa", []string{"aa", "aaa"}, true},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}, true},
		{"abcd", []string{"a", "abc", "b", "cd"}, true},
		{"abcd", []string{"ab", "abc", "cd", "def"}, true},
		{"abcd", []string{"ab", "abc", "de"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			actual := wordBreak(tt.s, tt.wordDict)
			if actual != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v; expected %v", tt.s, tt.wordDict, actual, tt.expected)
			}
		})
	}
}

func TestWordBreak_LongInput(t *testing.T) {
	longS := ""
	for i := 0; i < 1000; i++ {
		longS += "a"
	}
	longDict := []string{"a", "aa", "aaa", "aaaa"}
	expected := true

	actual := wordBreak(longS, longDict)
	if actual != expected {
		t.Errorf("wordBreak(longS of len %d, %v) = %v; expected %v", len(longS), longDict, actual, expected)
	}
}
