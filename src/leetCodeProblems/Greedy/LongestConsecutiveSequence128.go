package main

/*
- Leetcode - https://leetcode.com/problems/longest-consecutive-sequence
- Time - O(n)
- Space - O(n)
*/
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	out := 0

	for v := range m {
		if _, ok := m[v-1]; ok {
			continue
		}

		len := 0

		for {
			if _, ok := m[v+len]; !ok {
				break
			}
			len++
		}

		if len > out {
			out = len
		}
	}
	return out
}
