package main

import (
	"strings"
	"unicode"
)

/*
- LeetCode - https://leetcode.com/problems/valid-palindrome/
- Time - O(n)
- Space - O(1)
*/

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	s = strings.ToLower(s)

	for l < r {
		if !isAlphaNumeric(rune(s[l])) {
			l++
			continue
		}

		if !isAlphaNumeric(rune(s[r])) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
