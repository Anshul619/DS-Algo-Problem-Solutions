package main

/*
- LeetCode - https://leetcode.com/problems/next-greater-element-i/description/
- Time - O(m+n)
- Space - O(n)
*/
import "log"

type stack1 []int

func (s stack1) isEmpty() bool {
	return len(s) == 0
}

func (s stack1) peek() int {
	return s[len(s)-1]
}

func (s *stack1) pop() int {
	num := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return num
}

func (s *stack1) push(num int) {
	*s = append(*s, num)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	s := new(stack1)
	m := make(map[int]int)

	out := make([]int, len(nums1))

	for i := len(nums2) - 1; i >= 0; i-- {

		for !s.isEmpty() && s.peek() < nums2[i] {
			s.pop()
		}

		if !s.isEmpty() {
			m[nums2[i]] = s.peek()
		} else {
			m[nums2[i]] = -1
		}

		s.push(nums2[i])
	}

	for i, v := range nums1 {
		out[i] = m[v]
	}

	return out
}

func main() {
	log.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	log.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	log.Println(nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}))
}
