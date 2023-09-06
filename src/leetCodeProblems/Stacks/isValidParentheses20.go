package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/valid-parentheses/
- Time - O(n)
- Space - O(n)
*/

func isValid(s string) bool {

	stack := new(stackParenthesis)

	for _, v := range s {
		switch string(v) {
		case ")":
			if stack.isEmpty() || string(stack.peek()) != "(" {
				return false
			}
			stack.pop()
		case "}":
			if stack.isEmpty() || string(stack.peek()) != "{" {
				return false
			}
			stack.pop()
		case "]":
			if stack.isEmpty() || string(stack.peek()) != "[" {
				return false
			}
			stack.pop()
		default:
			stack.push(v)
		}
	}

	return stack.isEmpty()
}

func main() {
	log.Println(isValid("()[]{}"))
	log.Println(isValid("()"))
	log.Println(isValid("(]"))
	log.Println(isValid("[()]{}{[()()]()}"))
	log.Println(isValid("[(])"))
	log.Println(isValid("["))
}
