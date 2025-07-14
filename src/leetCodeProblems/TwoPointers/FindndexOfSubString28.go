package main

/*
- LeetCode - https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
- Time - O(n)
- Space - O(1)
*/

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
