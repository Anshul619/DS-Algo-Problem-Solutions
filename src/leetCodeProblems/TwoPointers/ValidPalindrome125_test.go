package main

/*
- LeetCode - https://leetcode.com/problems/valid-palindrome/
*/
import "testing"

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"test1", "A man, a plan, a canal: Panama", true},
		{"test2", "race a car", false},
		{"test3", " ", true},
		{"test4", "0P", false},
	}

	for _, v := range tests {
		if isPalindrome(v.input) != v.expected {
			t.Errorf("%s: failed", v.name)
		}
	}
}
