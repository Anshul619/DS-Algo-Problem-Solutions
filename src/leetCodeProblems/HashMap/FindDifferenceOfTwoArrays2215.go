package main

/*
- Leetcode - https://leetcode.com/problems/find-the-difference-of-two-arrays/description/
- Time - O(m+n)
- Space - O(m+n)
*/

func findDifference(nums1 []int, nums2 []int) [][]int {

	out := [][]int{[]int{}, []int{}}

	m1 := make(map[int]struct{})
	m2 := make(map[int]struct{})

	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		m2[v] = struct{}{}
	}

	for k, _ := range m2 {
		if _, ok := m1[k]; !ok {
			out[1] = append(out[1], k)
		}
	}

	for k, _ := range m1 {
		if _, ok := m2[k]; !ok {
			out[0] = append(out[0], k)
		}
	}

	return out
}
