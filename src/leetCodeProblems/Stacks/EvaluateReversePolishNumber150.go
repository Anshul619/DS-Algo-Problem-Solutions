package main

import "strconv"

/*
- LeetCode - https://leetcode.com/problems/evaluate-reverse-polish-notation
- Time - O(n)
- Space - O(n)
*/
type stackRPN []int

func (s *stackRPN) pop() int {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s *stackRPN) push(t int) {
	*s = append(*s, t)
}

func evalRPN(tokens []string) int {
	s := new(stackRPN)

	for _, v := range tokens {
		switch v {
		case "+", "-", "/", "*":
			n1 := s.pop()
			n2 := s.pop()
			switch v {
			case "+":
				s.push(n2 + n1)
			case "-":
				s.push(n2 - n1)
			case "/":
				s.push(n2 / n1)
			default:
				s.push(n2 * n1)
			}
		default:
			i, _ := strconv.Atoi(v)
			s.push(i)

		}
	}
	return s.pop()
}
