package main

/*
- Leetcode - https://leetcode.com/problems/valid-anagram/description/
- Time - O(n)
- Space - O(26)
*/
func isAnagram(s string, t string) bool {
	m := make([]int, 26)

	for _, v := range s {
		m[v-rune('a')]++
	}

	for _, v := range t {
		if m[v-rune('a')] <= 0 {
			return false
		}

		m[v-rune('a')]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
