package main

/**
- LeetCode - https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
- TimeComplexity - O(n)
- SpaceComplexity - O(n)
*/

type stack1 []rune

func (s *stack1) push(r rune) {
	*s = append(*s, r)
}

func (s *stack1) pop() rune {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s stack1) peek() rune {
	return (s)[len(s)-1]
}

func (s stack1) isEmpty() bool {
	return len(s) == 0
}

func (s stack1) get() []rune {
	return s
}

func removeDuplicates1(s string) string {
	st := new(stack1)

	for _, v := range s {
		if st.isEmpty() || st.peek() != v {
			st.push(v)
		} else {
			st.pop()
		}
	}
	return string(st.get())
}

// func main() {
// 	log.Println(removeDuplicates("abbaca"))
// 	log.Println(removeDuplicates("azxxzy"))
// }
