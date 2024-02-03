package main

/*
- LeetCode - https://leetcode.com/problems/contains-duplicate-ii
- Time - O(n)
- Space - O(n)
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if j, ok := m[v]; ok {
			if i-j <= k {
				return true
			}

		}
		m[v] = i
	}
	return false
}
