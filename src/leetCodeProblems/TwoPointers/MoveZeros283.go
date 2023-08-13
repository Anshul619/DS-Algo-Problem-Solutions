package main

/*
- LeetCode - https://leetcode.com/problems/move-zeroes/
- Time - O(n)
- Space - O(1)
*/

func moveZeroes(nums []int) {
	start, iterator := 0, 0

	for iterator < len(nums) {

		if nums[iterator] == 0 {
			iterator++
			continue
		}
		nums[start] = nums[iterator]
		start++
		iterator++
	}

	for i := start; i < len(nums); i++ {
		nums[i] = 0
	}
}

// func main() {
// 	nums1 := []int{0, 1, 0, 3, 12}
// 	moveZeroes(nums1)

// 	log.Println(nums1)

// 	nums1 = []int{0}

// 	moveZeroes(nums1)

// 	log.Println(nums1)
// }
