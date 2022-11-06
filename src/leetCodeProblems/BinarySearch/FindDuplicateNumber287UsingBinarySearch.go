package main

/*
- LeetCode - https://leetcode.com/problems/find-the-duplicate-number/
- Time Complexity - O(nlogn)
- Space Complexity - O(1)
*/

import "log"

func findDuplicate(nums []int) int {

	low := 0
	high := len(nums) - 1

	duplicate := -1

	for low <= high {

		cur := low + (high-low)/2
		count := 0

		for _, v := range nums {
			if v <= cur {
				count++
			}
		}

		log.Println(cur, count)

		if count > cur {
			duplicate = cur
			high = cur - 1
		} else {
			low = cur + 1
		}
	}

	return duplicate
}

func main() {

	nums := []int{1, 3, 4, 2, 2}

	//nums := []int{3, 1, 3, 4, 2}

	// nums := []int{2, 2, 2, 2, 2}

	log.Println(findDuplicate(nums))
}
