package main

/*
- LeetCode - https://leetcode.com/problems/contains-duplicate/
- Time - O(n)
- Space - O(n)
*/

func containsDuplicate(nums []int) bool {
	// Create an empty map to track seen integers
	m := make(map[int]bool)

	// Iterate over all numbers in nums
	for _, v := range nums {
		// If v is already in the map, we found a duplicate
		if _, ok := m[v]; ok {
			return true
		}
		// Otherwise, mark v as seen
		m[v] = true
	}
	// No duplicates found
	return false
}
