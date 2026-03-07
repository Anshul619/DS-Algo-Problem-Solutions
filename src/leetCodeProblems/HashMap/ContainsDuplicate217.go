package main

/*
- LeetCode - https://leetcode.com/problems/contains-duplicate/
- Time - O(n)
- Space - O(n)
*/

func containsDuplicate(nums []int) bool {
	// Create an empty map to track seen integers
	seen := make(map[int]struct{})

	// Iterate over all numbers in nums
	for _, v := range nums {
		// If v is already in the map, we found a duplicate
		if _, ok := seen[v]; ok {
			return true
		}
		// Otherwise, mark v as seen
		seen[v] = struct{}{}
	}
	// No duplicates found
	return false
}
