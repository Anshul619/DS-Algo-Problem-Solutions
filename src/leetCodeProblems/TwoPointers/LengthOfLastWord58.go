package main

/*
- LeetCode - https://leetcode.com/problems/length-of-last-word/
- Time - O(n)
- Space - O(1)
*/
func lengthOfLastWord(s string) int {
	start, end, endFound := -1, -1, false

	for i := len(s) - 1; i >= 0; i-- {
		if !endFound && string(s[i]) != " " {
			end = i
			endFound = true
		}

		if endFound && string(s[i]) == " " {
			start = i
			break
		}
	}
	return end - start
}
