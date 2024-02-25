package main

/*
- Leetcode - https://leetcode.com/problems/word-pattern/description
- Time - O(n)
- Space - O(n)
*/
func findNextWord(s string, i int) (string, int) {
	out := ""

	for i < len(s) && s[i] != byte(' ') {
		out += string(s[i])
		i++

	}

	// one more increase fo space
	i++
	return out, i
}
func wordPattern(pattern string, s string) bool {
	m := make(map[rune]string)
	mr := make(map[string]rune)

	i := 0

	for _, v := range pattern {
		var w string
		w, i = findNextWord(s, i)

		if w == "" {
			return false
		}

		if p1, ok := mr[w]; ok && p1 != v {
			return false
		}
		if p, ok := m[v]; ok && p != w {
			return false
		}
		m[v] = w
		mr[w] = v
	}

	return i >= len(s)
}
