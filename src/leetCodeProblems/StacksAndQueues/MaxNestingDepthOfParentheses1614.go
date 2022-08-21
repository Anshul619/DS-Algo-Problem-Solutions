package main

/*
- LeetCode - https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
*/
import "log"

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}
func (s *Stack) Len() int {
	return len(*s)
}
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		elem := (*s)[index]
		*s = (*s)[:index]
		return elem, true
	}
}

func maxDepth(s string) int {

	stack := new(Stack)
	max := 0

	for _, v := range s {
		if v == '(' {
			stack.Push("(")
		} else if v == ')' {

			if stackLen := stack.Len(); max < stackLen {
				max = stackLen
			}

			stack.Pop()
		}
	}

	return max
}

func main() {

	//input := "(1+(2*3)+((8)/4))+1"
	input := "(1)+((2))+(((3)))"
	log.Println(maxDepth(input))
}
