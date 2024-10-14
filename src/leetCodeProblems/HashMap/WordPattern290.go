package main

/*
- Leetcode - https://leetcode.com/problems/word-pattern/description
- Time - O(n)
- Space - O(m+n) // (where m is length of pattern and n is length of string)
*/
func findNextWord(i *int, s string) string {
	old := *i

	for *i < len(s) && s[*i] != byte(' ') {
		*i++
	}

	if old == *i {
		return ""
	}

	// one more increase fo space
	*i++
	return string(s[old : *i-1])
}

func wordPattern(pattern string, s string) bool {
	m := make(map[rune]string)
	mr := make(map[string]rune)

	i := 0

	for _, v := range pattern {
		w := findNextWord(&i, s)

		if w == "" {
			return false
		}

		if _, ok := m[v]; ok && m[v] != w {
			return false
		}

		if _, ok := mr[w]; ok && mr[w] != v {
			return false
		}

		m[v] = w
		mr[w] = v
	}

	return i >= len(s)
}
