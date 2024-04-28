package main

/*
- LeetCode - https://leetcode.com/problems/evaluate-reverse-polish-notation
- Time - O(n)
- Space - O(n)
*/
import (
	"strconv"
)

type stackRPN []string

func (s *stackRPN) push(str string) {
	*s = append(*s, str)
}

func (s *stackRPN) pop() string {
	str := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return str
}

func (s *stackRPN) pop2Times() (string, string) {
	str := (*s)[len(*s)-2:]
	*s = (*s)[:len(*s)-2]
	return str[0], str[1]
}

func evalRPN(tokens []string) int {

	type fn func(int, int) int

	s := new(stackRPN)

	m := map[string]fn{
		"+": func(a int, b int) int { return a + b },
		"-": func(a int, b int) int { return a - b },
		"*": func(a int, b int) int { return a * b },
		"/": func(a int, b int) int { return a / b },
	}

	for _, v := range tokens {

		if fn, ok := m[v]; ok {

			n1, n2 := s.pop2Times()

			n3, _ := strconv.Atoi(n1)
			n4, _ := strconv.Atoi(n2)

			v = strconv.Itoa(fn(n3, n4))

		}
		s.push(v)
	}

	out, _ := strconv.Atoi(s.pop())

	return out
}
