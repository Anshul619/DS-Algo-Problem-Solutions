package main

/*
- LeetCode - https://leetcode.com/problems/subarray-sum-equals-k/
- Time - O(n)
- Extra Space - O(n)
*/

func subarraySum(nums []int, k int) int {
	m := make(map[int]int) // stores how many times, each prefix sum has occurred

	cur := 0 // tracks running total
	out := 0
	m[0] = 1

	for _, v := range nums {
		cur += v

		// If cur - k has occurred before, then there exists a subarray ending at the current index that sums to k.
		if count, ok := m[cur-k]; ok {
			out += count
		}

		m[cur]++
	}
	return out
}
