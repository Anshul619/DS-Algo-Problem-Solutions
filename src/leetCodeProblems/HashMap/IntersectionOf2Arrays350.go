package main

/*
- LeetCode - https://leetcode.com/problems/intersection-of-two-arrays-ii/
- Space - O(n)
- Time - O(n)
*/

func pushToMap(arr []int, m map[int]int) {
	for _, v := range arr {
		m[v]++
	}
}

func checkMap(arr []int, m map[int]int) []int {
	out := []int{}

	for _, v := range arr {
		if count, ok := m[v]; ok && count > 0 {
			out = append(out, v)
			m[v]--
		}
	}

	return out
}

func intersect(nums1 []int, nums2 []int) []int {

	m := make(map[int]int)

	if len(nums1) > len(nums2) {
		pushToMap(nums1, m)
		return checkMap(nums2, m)

	} else {
		pushToMap(nums2, m)
		return checkMap(nums1, m)
	}
}
