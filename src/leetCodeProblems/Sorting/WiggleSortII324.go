package main

/*
- LeetCode - hhttps://leetcode.com/problems/wiggle-sort-ii/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"sort"
)

func wiggleSort(nums []int) {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	mid := (len(nums) - 1) / 2
	right := len(nums) - 1

	for i := range nums {
		if i%2 == 0 {
			nums[i] = sorted[mid]
			mid--
		} else {
			nums[i] = sorted[right]
			right--
		}
	}
}

// func main() {

// 	nums := []int{1, 5, 1, 1, 6, 4}

// 	wiggleSort(nums)
// 	log.Println(nums)
// }
