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

// func main() {
// 	log.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
// 	log.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
// 	log.Println(intersect([]int{3, 1, 2}, []int{1, 1}))
// }
