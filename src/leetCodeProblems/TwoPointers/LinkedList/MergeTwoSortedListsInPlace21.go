package main

/*
- Leetcode - https://leetcode.com/problems/merge-sorted-array/description/
- Time - O(m+n)
- Space - O(1)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	out := m + n - 1
	m--
	n--

	for m >= 0 && n >= 0 {
		if nums2[n] >= nums1[m] {
			nums1[out] = nums2[n]
			out--
			n--
		} else {
			nums1[out] = nums1[m]
			out--
			m--
		}
	}

	for n >= 0 {
		nums1[out] = nums2[n]
		out--
		n--
	}
}
