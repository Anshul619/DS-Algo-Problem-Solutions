package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/kth-missing-positive-number/description/
- Time - O(n)
- Space - O(1)
*/

func findKthPositive1(arr []int, k int) int {

	diff := 0

	for i, v := range arr {
		diff = v - 1

		if i > 0 {
			diff -= arr[i-1]
		}

		switch {
		case diff == k:
			return v - 1
		case diff > k:
			if i == 0 {
				return k
			}
			return arr[i-1] + k
		default:
			k -= diff
		}
	}

	return arr[len(arr)-1] + k
}

func main() {
	// log.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
	// log.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
	// log.Println(findKthPositive([]int{2}, 1))
	log.Println(findKthPositive1([]int{7, 13, 21, 25, 29, 32, 38, 45}, 4))
}
