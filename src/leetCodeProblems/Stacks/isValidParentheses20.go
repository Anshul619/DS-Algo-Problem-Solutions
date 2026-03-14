package main

/*
- LeetCode - https://leetcode.com/problems/valid-parentheses/
- Time - O(n)
- Space - O(n)
*/

func isValid(s string) bool {
	st := make([]rune, 0, len(s))

	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, v := range s {
		if open, ok := m[v]; ok {
			if len(st) == 0 || st[len(st)-1] != open {
				return false
			}
			st = st[:len(st)-1]
		} else {
			st = append(st, v)
		}
	}
	return len(st) == 0
}
