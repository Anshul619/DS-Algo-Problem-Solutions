package main

/*
- Leetcode - https://leetcode.com/problems/h-index/description/
- Time - O(n)
- Space - O(n)
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func hIndex(citations []int) int {
	counter := make([]int, len(citations)+1)

	for _, v := range citations {
		counter[min(len(citations), v)]++
	}

	max := len(citations)

	for i := counter[len(citations)]; max > i; i += counter[max] {
		max--
	}

	return max
}
