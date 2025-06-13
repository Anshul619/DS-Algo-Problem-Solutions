package main

/*
- Leetcode - https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/
- Time - O(n)
- Space - O(1)
*/

func longestSubstringWithAtMostKUniqueCharacters340(s string, k int) int {
	if k == 0 {
		return 0
	}

	// tracks count of characters in map
	m := make(map[rune]int)

	out := 0
	start := 0

	for i, v := range s {
		m[v]++

		if len(m) > k {
			startChar := rune(s[start])
			m[startChar]--
			if m[startChar] == 0 {
				delete(m, startChar)
			}
			start++
		}

		if i-start+1 > out {
			out = i - start + 1
		}
	}
	return out
}
