package main

/*
- LeetCode - https://leetcode.com/problems/intersection-of-two-arrays-ii/
- Space - O(n)
- Time - O(n)
*/

func intersect2(nums1 []int, nums2 []int) []int {

	m := make(map[int]int)

	for _, v := range nums1 {
		m[v]++
	}

	out := []int{}
	for _, v := range nums2 {
		if m[v] > 0 {
			out = append(out, v)
			m[v]--
		}
	}
	return out
}
