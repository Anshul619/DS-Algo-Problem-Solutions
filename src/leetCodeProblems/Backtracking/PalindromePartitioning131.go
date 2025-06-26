package main

/*
- Leetcode - https://leetcode.com/problems/palindrome-partitioning/
- Time - O(N^2)
- Space - O(N^2)
*/
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func dfs(s string, cur *[]string, out *[][]string, curIndex int) {
	if curIndex >= len(s) {
		*out = append(*out, append([]string{}, *cur...))
		return
	}

	for i := curIndex; i < len(s); i++ {
		subStr := s[curIndex : i+1]
		if isPalindrome(subStr) {
			*cur = append(*cur, subStr)
			dfs(s, cur, out, i+1)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}

func partition(s string) [][]string {
	out := [][]string{}
	cur := []string{}
	dfs(s, &cur, &out, 0)
	return out
}
