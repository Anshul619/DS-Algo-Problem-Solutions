package main

/*
- Leetcode - https://leetcode.com/problems/find-the-distinct-difference-array/description/
- Time - O(n)
- Space - O(n)
*/
func distinctDifferenceArray(nums []int) []int {
	r, l := make(map[int]int), make(map[int]int)

	out := make([]int, len(nums))

	for _, v := range nums {
		r[v]++
	}

	for i, v := range nums {
		l[v]++
		r[v]--
		if r[v] == 0 {
			delete(r, v)
		}

		out[i] = len(l) - len(r)
	}
	return out
}
