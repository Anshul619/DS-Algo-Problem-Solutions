package main

/*
- LeetCode - https://leetcode.com/problems/valid-parentheses/
- Time - O(n)
- Space - O(n)
*/

type stackP []rune

func (s *stackP) push(r rune) {
	*s = append(*s, r)
}

func (s *stackP) pop() rune {
	if s.isEmpty() {
		return rune(0)
	}

	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s stackP) isEmpty() bool {
	return len(s) == 0
}

func isValid(s string) bool {
	st := new(stackP)

	for _, v := range s {
		switch string(v) {
		case ")":
			e := st.pop()
			if string(e) != "(" {
				return false
			}
		case "}":
			e := st.pop()
			if string(e) != "{" {
				return false
			}
		case "]":
			e := st.pop()
			if string(e) != "[" {
				return false
			}
		default:
			st.push(v)
		}
	}

	return st.isEmpty()
}
