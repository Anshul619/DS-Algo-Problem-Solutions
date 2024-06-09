package main

/**
- LeetCode - https://leetcode.com/problems/daily-temperatures
- Time - O(n)
- Space - O(n)
*/

type Temp struct {
	Val int
	Day int
}

type tStack []Temp

func (s *tStack) push(t Temp) {
	*s = append(*s, t)
}

func (s *tStack) pop() Temp {
	t := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t
}

func (s tStack) isEmpty() bool {
	return len(s) == 0
}

func (s tStack) peek() Temp {
	return s[len(s)-1]
}

func dailyTemperatures(temperatures []int) []int {
	out := make([]int, len(temperatures))

	s := new(tStack)

	for i := len(temperatures) - 1; i >= 0; i-- {
		for !s.isEmpty() &&
			s.peek().Val <= temperatures[i] {
			s.pop()
		}

		if !s.isEmpty() {
			out[i] = s.peek().Day - i
		}

		s.push(Temp{
			Val: temperatures[i],
			Day: i,
		})
	}
	return out
}
