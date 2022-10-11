package main

/*
- LeetCode - https://leetcode.com/problems/single-element-in-a-sorted-array/
- TimeComplexity - O(logn)
- SpaceComplexity - O(1)
*/
import "log"

func util(nums []int, start int, end int) int {

	if start < 0 || end >= len(nums) {
		return 0
	}

	if start == end {
		return nums[start]
	}

	mid := start + (end-start)/2

	if nums[mid] == nums[mid-1] {
		leftLength := mid - 1 - start

		if leftLength%2 == 0 {
			start = mid + 1
		} else {
			end = mid - 2
		}

	} else if nums[mid] == nums[mid+1] {
		leftLength := mid - start

		if leftLength%2 == 0 {
			start = mid + 2
		} else {
			end = mid - 1
		}

	} else {
		return nums[mid]
	}

	return util(nums, start, end)
}
func singleNonDuplicate(nums []int) int {
	return util(nums, 0, len(nums)-1)
}

func main() {
	//nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	nums := []int{3, 3, 7, 7, 10, 11, 11}
	log.Println(singleNonDuplicate(nums))
}
