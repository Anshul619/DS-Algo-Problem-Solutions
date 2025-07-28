package main

/*
- Leetcode - https://leetcode.com/problems/longest-consecutive-sequence
- Time - O(n)
- Space - O(n)
*/
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		m[v] = true
	}

	out := 0

	for _, v := range nums {
		if ok := m[v-1]; !ok {
			j := 0

			for m[v+j] {
				j++
			}

			if j > out {
				out = j
			}

			if j > len(nums)/2 {
				break
			}
		}
	}
	return out
}
