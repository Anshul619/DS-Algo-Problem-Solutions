package main

/*
- LeetCode - https://leetcode.com/problems/intersection-of-two-arrays
- Space - O(n)
- Time - O(n)
*/

func pushToMap1(arr []int, m map[int]bool) {
	for _, v := range arr {
		m[v] = false
	}
}

func checkMap1(arr []int, m map[int]bool) []int {
	out := []int{}

	for _, v := range arr {
		if alreadyChecked, ok := m[v]; ok && !alreadyChecked {
			out = append(out, v)
			m[v] = true
		}
	}

	return out
}

func intersection(nums1 []int, nums2 []int) []int {

	m := make(map[int]bool)

	if len(nums1) > len(nums2) {
		pushToMap1(nums1, m)
		return checkMap1(nums2, m)

	} else {
		pushToMap1(nums2, m)
		return checkMap1(nums1, m)
	}
}

// func main() {
// 	log.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
// }
