package main

/*
- Leetcode - https://leetcode.com/problems/longest-consecutive-sequence
- Time - O(n)
- Space - O(n)
*/
func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	out := 0

	for _, v := range nums {
		m[v] = true
	}

	for _, v := range nums {
		if _, ok := m[v-1]; !ok {

			j := 0

			for {
				if _, ok := m[v+j]; ok {
					j++
				} else {
					break
				}
			}

			out = max(out, j)

			if j > len(nums)/2 {
				break
			}
		}
	}
	return out
}
