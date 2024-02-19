package main

/*
- Leetcode - https://leetcode.com/problems/merge-intervals/description/
- Time - O(nlogn)
- Space - O()
*/

import (
	"sort"
)

type stack [][]int

func (s *stack) push(num []int) {
	*s = append(*s, num)
}

func (s *stack) pop() []int {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s stack) peek() []int {
	return s[len(s)-1]
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeUsingStack(intervals [][]int) [][]int {
	out := [][]int{}

	s := new(stack)

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	if len(intervals) > 0 {
		s.push(intervals[0])
	}

	for i := 1; i < len(intervals); i++ {
		last := s.peek()

		if last[1] >= intervals[i][0] {
			s.pop()
			s.push([]int{last[0], max(intervals[i][1], last[1])})
		} else {
			s.push(intervals[i])
		}
	}

	for !s.isEmpty() {
		out = append(out, s.pop())
	}
	return out
}

// func main() {
// 	//log.Println(merge([][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}))
// 	log.Println(merge([][]int{{1, 4}, {4, 5}}))
// }
