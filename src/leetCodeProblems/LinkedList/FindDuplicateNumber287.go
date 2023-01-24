package main

/*
- LeetCode - https://leetcode.com/problems/find-the-duplicate-number/
- Time Complexity - O(logn)
- Space Complexity - O(1)
- Remarks - Floyd Algo
*/

func findDuplicate(nums []int) int {

	slow := nums[0]
	fast := nums[0]

	// Find intersection point
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			break
		}

	}

	slow = nums[0]

	// Find the point where cycle starts
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}

// func main() {

// 	//nums := []int{1, 3, 4, 2, 2}

// 	nums := []int{3, 1, 3, 4, 2}

// 	//nums := []int{3, 1, 3, 4, 2}

// 	// nums := []int{2, 2, 2, 2, 2}

// 	log.Println(findDuplicate(nums))
// }
