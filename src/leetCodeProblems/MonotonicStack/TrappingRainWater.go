package main

/*
- Leetcode - https://leetcode.com/problems/trapping-rain-water/description/
- time - O(n)
- space - O(n)
*/
import "log"

type stack []int

func (s *stack) push(n int) {
	*s = append(*s, n)
}

func (s *stack) pop() int {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

func (s stack) peek() int {
	return s[len(s)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func trap(height []int) int {

	s := new(stack)
	total := 0

	for i, v := range height {

		log.Println(s)
		for !s.isEmpty() && height[s.peek()] < v {
			popped_i := s.pop()
			if s.isEmpty() {
				break
			}

			h := min(height[s.peek()], v) - height[popped_i]
			len := i - s.peek() - 1
			total += h * len
		}
		s.push(i)
	}

	return total
}

func main() {
	log.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
