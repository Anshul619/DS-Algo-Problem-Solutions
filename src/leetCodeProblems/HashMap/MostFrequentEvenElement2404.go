package main

/*
- LeetCode - https://leetcode.com/problems/most-frequent-even-element/description/
- Time - O(n)
- Space - O(n)
*/
func mostFrequentEven(nums []int) int {

	m := make(map[int]int)

	for _, v := range nums {
		if v%2 == 0 {
			m[v]++
		}
	}

	out, count := -1, 0

	for i, v := range m {
		switch {
		case v > count:
			out = i
			count = v
		case v == count:
			if i < out {
				out = i
			}
		}
	}
	return out
}
