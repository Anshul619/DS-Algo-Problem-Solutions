package main

/*
- LeetCode - https://leetcode.com/problems/contains-duplicate/
- Time - O(n)
- Space - O(n)
*/

func containsDuplicate(nums []int) bool {

	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
