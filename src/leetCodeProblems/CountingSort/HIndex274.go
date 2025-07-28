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
	counter := make([]int, len(citations)+1) // counter[i] = number of papers with exactly i citations
	// All papers with >n citations are grouped into counter[n]

	// Step 1: Count how many papers have each citation count (capped at n)
	for _, v := range citations {
		counter[min(len(citations), v)]++
	}

	// Step 2: Start from the maximum possible H-index (which is n)
	h := len(citations)

	// papersWithAtLeastH keeps track of the total number of papers with at least `h` citations
	// Initially, it's the number of papers with >= n citations
	papersWithAtLeastH := counter[h]

	// Step 3: Decrease `h` until we find a valid H-index
	// We want to find the largest `h` where there are at least `h` papers with >= h citations
	for h > papersWithAtLeastH {
		h--                              // Try a smaller h-index
		papersWithAtLeastH += counter[h] // Add the number of papers that have exactly `h` citations
	}

	// When we exit the loop, h <= papersWithAtLeastH, meaning
	// there are at least `h` papers with >= `h` citations
	return h
}
