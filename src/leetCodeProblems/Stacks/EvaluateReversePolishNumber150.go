package main

import "strconv"

/*
- LeetCode - https://leetcode.com/problems/evaluate-reverse-polish-notation
- Time - O(n)
- Space - O(n)
*/
func evalRPN(tokens []string) int {
	st := []int{}

	for _, v := range tokens {
		switch v {
		case "-", "/", "*", "+":
			b := st[len(st)-1]
			a := st[len(st)-2]
			st = st[:len(st)-2]

			switch v {
			case "+":
				st = append(st, a+b)
			case "-":
				st = append(st, a-b)
			case "/":
				st = append(st, a/b)
			default:
				st = append(st, a*b)
			}
		default:
			n, _ := strconv.Atoi(v)
			st = append(st, n)
		}
	}
	return st[0]
}
