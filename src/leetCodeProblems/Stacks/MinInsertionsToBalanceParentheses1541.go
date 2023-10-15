package main

/*
- Leetcode - https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string/description/
- Not-Completed
- Time - O(n)
- Space - O(n)
*/
import "log"

type Parenthesis struct {
	p rune
	i int
}

type stackParenthesis1 []Parenthesis

func (s *stackParenthesis1) pop() Parenthesis {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *stackParenthesis1) push(r Parenthesis) {
	*s = append(*s, r)
}

func (s stackParenthesis1) isEmpty() bool {
	return len(s) == 0
}

func (s stackParenthesis1) peek() Parenthesis {
	return s[len(s)-1]
}

func (s stackParenthesis1) len() int {
	return len(s)
}

func minInsertions(s string) int {
	st := new(stackParenthesis1)
	out := 0

	i := 0

	for i < len(s) {
		if string(rune(s[i])) == ")" &&
			i+1 < len(s) &&
			string(rune(s[i+1])) == ")" &&
			!st.isEmpty() && string(st.peek().p) == "(" {
			st.pop()
			i = i + 2
			continue
		}
		st.push(Parenthesis{rune(s[i]), i})
		i++
	}

	log.Println(st)
	for !st.isEmpty() {
		if string(st.peek().p) == ")" {
			c := st.pop()

			if !st.isEmpty() {
				if string(st.peek().p) == ")" && st.peek().i == c.i+1 {
					st.pop()

					if !st.isEmpty() {
						if string(st.peek().p) != "(" {
							out++
						} else {
							st.pop()
						}
					} else {
						out++
					}
				} else {
					//st.pop()
					out += 2
				}
			} else {
				out += 2
			}
		} else {
			st.pop()
			out += 2
		}
	}

	log.Println(out)
	return out
}
