package main

/*
- LeetCode - https://leetcode.com/problems/find-the-duplicate-number/
- Time Complexity - O(n)
- Space Compleixty - O(1)
*/

func findDuplicate(nums []int) int {

	for nums[0] != nums[nums[0]] {
		temp := nums[nums[0]]
		nums[nums[0]] = nums[0]
		nums[0] = temp
	}

	return nums[0]
}

// func main() {

// 	nums := []int{1, 3, 4, 2, 2}

// 	//nums := []int{3, 1, 3, 4, 2}

// 	// nums := []int{2, 2, 2, 2, 2}

// 	log.Println(findDuplicate(nums))
// }
