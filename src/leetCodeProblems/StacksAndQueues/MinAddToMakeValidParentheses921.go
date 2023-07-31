package main

/*
- Leetcode - https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/description/
- Space - O(n)
- Time - O(n)
*/

type stackParenthesis []rune

func (s *stackParenthesis) pop() rune {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *stackParenthesis) push(r rune) {
	*s = append(*s, r)
}

func (s stackParenthesis) isEmpty() bool {
	return len(s) == 0
}

func (s stackParenthesis) peek() rune {
	return s[len(s)-1]
}

func (s stackParenthesis) len() int {
	return len(s)
}

func minAddToMakeValid(s string) int {
	st := new(stackParenthesis)

	for _, r := range s {
		if st.isEmpty() {
			st.push(r)
		} else {

			if string(r) == ")" && string(st.peek()) == "(" {
				st.pop()
			} else {
				st.push(r)
			}
		}
	}

	return st.len()
}

// func main() {
// 	log.Println(minAddToMakeValid("())"))
// }
