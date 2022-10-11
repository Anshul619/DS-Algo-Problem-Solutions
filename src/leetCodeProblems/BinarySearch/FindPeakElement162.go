package main

/*
- LeetCode - https://leetcode.com/problems/find-peak-element/submissions/
*/

func binarySearch(nums []int, start int, end int) (int, bool) {

	mid := start + (end-start)/2

	if start < 0 || end > len(nums) || mid < 0 {
		return -1, false
	}

	if (mid == 0 || nums[mid] >= nums[mid-1]) && (mid == len(nums)-1 || nums[mid] >= nums[mid+1]) {
		return mid, true
	} else if mid > 0 && nums[mid-1] >= nums[mid] {
		return binarySearch(nums, start, mid-1)
	} else {
		return binarySearch(nums, mid+1, end)
	}
}

func findPeakElement(nums []int) int {

	if len(nums) == 1 {
		return 0
	}
	peak, _ := binarySearch(nums, 0, len(nums))
	return peak
}

// func main() {

// 	//input := []int{1, 2, 3, 1}
// 	input := []int{1, 2, 3}
// 	log.Println(findPeakElement(input))
// }
