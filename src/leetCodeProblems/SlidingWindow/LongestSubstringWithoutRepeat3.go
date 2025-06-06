package main

/*
- Leet Code - https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
- Space - O(n)
- Time - O(n)
*/

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	out := 0

	// Tracks the beginning of the current window of unique characters.
	start := 0

	for i, v := range s {

		if _, ok := m[v]; ok &&
			m[v] >= start { // Ensures you don’t move the window backward when seeing old duplicates.
			start = m[v] + 1
		}

		if i-start+1 > out {
			out = i - start + 1
		}

		// Keeps the latest index of each character.
		m[v] = i
	}
	return out
}
