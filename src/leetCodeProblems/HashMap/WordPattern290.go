package main

import "strings"

/*
- Leetcode - https://leetcode.com/problems/word-pattern/description
- Time - O(n)
- Space - O(m+n) // (where m is length of pattern and n is length of string)
*/
func wordPattern(pattern string, s string) bool {
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)

	stringSlice := strings.Split(s, " ")

	if len(pattern) != len(stringSlice) {
		return false
	}

	for i, v := range pattern {
		m1[byte(v)] = stringSlice[i]
		m2[stringSlice[i]] = byte(v)
	}

	for i, v := range stringSlice {
		if m2[v] != pattern[i] || m1[pattern[i]] != v {
			return false
		}
	}

	return true
}
