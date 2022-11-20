package main

/*
- LeetCode - https://leetcode.com/problems/remove-duplicates-from-sorted-array/
- Time - O(n)
- Space - O(1)
*/
import "log"

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	firstPointer := 0
	secondPointer := 1
	lastElement := nums[firstPointer]

	for secondPointer < len(nums) {

		if nums[secondPointer] != lastElement {
			firstPointer++

			temp := nums[secondPointer]
			nums[secondPointer] = nums[firstPointer]
			nums[firstPointer] = temp

			lastElement = nums[firstPointer]
		}

		secondPointer++
	}

	return firstPointer + 1
}

func main() {
	//nums := []int{1, 1, 2}

	//nums := []int{1}

	//nums := []int{1, 2}

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	log.Println(removeDuplicates(nums))
}
