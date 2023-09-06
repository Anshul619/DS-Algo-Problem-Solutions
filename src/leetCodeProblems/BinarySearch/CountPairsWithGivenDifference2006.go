package main

import (
	"log"
	"sort"
)

/*
- LeetCode - https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/description/
- Time - O(nlogn)
- Space - O(1)
*/

func binarySearchWithDuplicatesHandling(nums []int, start, search int) (int, bool) {

	end := len(nums) - 1
	out := len(nums)

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] >= search {
			out = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return out, false
}

func countKDifference(nums []int, k int) int {
	sort.Ints(nums)
	out := 0

	for i, v := range nums {
		if searchIndex, _ := binarySearchWithDuplicatesHandling(nums, i+1, v+k); searchIndex != len(nums) {
			incrementedSearchIndex, _ := binarySearchWithDuplicatesHandling(nums, i+1, v+k+1)
			out += incrementedSearchIndex - searchIndex
		}
	}

	return out
}

func main() {

	//log.Println(countKDifference([]int{1, 2, 2, 1}, 1))
	log.Println(countKDifference([]int{10, 2, 10, 9, 1, 6, 8, 9, 2, 8}, 5))
}
