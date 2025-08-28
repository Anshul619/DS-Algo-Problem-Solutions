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

func isAlphanumeric(b byte) bool {
	if unicode.IsLetter(rune(b)) || unicode.IsNumber(rune(b)) {
		return true
	}

	return false
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	s = strings.ToLower(s)

	for left < right {
		if left == right {
			return false
		}
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
