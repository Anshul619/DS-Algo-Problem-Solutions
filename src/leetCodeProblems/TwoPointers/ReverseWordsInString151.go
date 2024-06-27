package main

/*
- Leetcode - https://leetcode.com/problems/reverse-words-in-a-string/description/
- Time - O(n)
- Space - O(1)
*/

func reverseWords(s string) string {
	lastIndex := -1
	out := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {

			if lastIndex == -1 {
				lastIndex = i
			}
			continue
		}

		if lastIndex != -1 {

			if len(out) != 0 {
				out += " "
			}
			out += s[i+1 : lastIndex+1]
			lastIndex = -1
		}
	}

	if lastIndex != -1 {

		if len(out) != 0 {
			out += " "
		}
		out += s[0 : lastIndex+1]
	}
	return out
}
