package main

/*
- Leetcode - https://leetcode.com/problems/merge-sorted-array
- Time - O(m+n)
- Space - O(1)
*/
func MergeArray(nums1 []int, m int, nums2 []int, n int) {
	out := len(nums1) - 1
	m--
	n--

	for m > -1 && n > -1 {
		if nums1[m] > nums2[n] {
			nums1[out] = nums1[m]
			m--
		} else {
			nums1[out] = nums2[n]
			n--
		}
		out--
	}

	for n > -1 {
		nums1[out] = nums2[n]
		n--
		out--
	}
}
