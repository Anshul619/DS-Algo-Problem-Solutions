package main

/*
- LeetCode - https://leetcode.com/problems/find-pivot-index/
- Time - O(n)
- Extra Space - O(1)
*/

func pivotIndex(nums []int) int {
	left, total := 0, 0

	for _, v := range nums {
		total += v
	}

	for i, v := range nums {
		if left == total-left-v {
			return i
		}

		left += v
	}

	return -1
}

// func main() {
// 	log.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
// }
