package main

/*
- LeetCode - https://leetcode.com/problems/sort-colors/submissions/
- Time - O(n)
- Space - O(1)
*/

func swap(index1 int, index2 int, nums []int) {
	temp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = temp
}

func sortColors(nums []int) {

	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		if nums[mid] == 0 {
			swap(low, mid, nums)
			low++
			mid++
		} else if nums[mid] == 2 {
			swap(mid, high, nums)
			high--
		} else { //nums[mid] == 1
			mid++
		}

	}
}

// func main() {
// 	//nums := []int{2, 0, 2, 1, 1, 0}

// 	nums := []int{1, 2, 0}
// 	sortColors(nums)
// 	log.Println(nums)
// }
