package main

/*
- Leetcode - https://leetcode.com/problems/longest-common-prefix/description/
- Time - O(n)
- Space - O(1)
*/
func findCommonPrefix(str1, str2 string) string {
	i := 0

	for i < len(str1) && i < len(str2) {
		if str1[i] != str2[i] {
			break
		}

		i++
	}
	return str1[0:i]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	out := strs[0]

	for i := 1; i < len(strs); i++ {
		out = findCommonPrefix(out, strs[i])

		if out == "" {
			break
		}
	}
	return out
}
