package main

/*
- LeetCode - https://leetcode.com/problems/two-sum/description/
- Time - O(n)
- Space - O(n)
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		m[v] = i
	}

	return []int{}
}
