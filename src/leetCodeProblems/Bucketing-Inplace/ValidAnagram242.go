package main

/*
- LeetCode - https://leetcode.com/problems/valid-anagram/description/
- Time complexity: O(n)
- Space complexity: O(26) = O(1)
*/

/*
Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
- Have to use map[rune]int, instead of make([]int, 26)
- Time - O(n)
- Space - O(n)
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)

	for _, v := range s {
		m[v-'a']++
	}

	for _, v := range t {
		m[v-'a']--
		if m[v-'a'] < 0 {
			return false
		}
	}

	return true
}
