package main

/*
- Leetcode - https://leetcode.com/problems/generate-parentheses/
- Time - O(4ⁿ / √n)
- Space - O(n)
*/
func generateParenthesisUtils(n int, open int, close int, str string, out *[]string) {
	if open == n && close == n {
		*out = append(*out, str)
		return
	}

	if open < n {
		generateParenthesisUtils(n, open+1, close, str+"(", out)
	}

	if close < open {
		generateParenthesisUtils(n, open, close+1, str+")", out)
	}
}
func generateParenthesis(n int) []string {
	out := []string{}

	generateParenthesisUtils(n, 0, 0, "", &out)
	return out
}
