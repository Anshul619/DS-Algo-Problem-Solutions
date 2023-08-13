package main

/*
- LeetCode - https://leetcode.com/problems/remove-element/
- Time - O(n)
- Space - O(1)
*/

func removeElement(nums []int, val int) int {
	start, end := 0, len(nums)-1

	for start <= end {

		if nums[end] == val {
			end--
			continue
		}

		if nums[start] == val {
			nums[start] = nums[end]
			end--
		}
		start++
	}
	return end + 1
}

// func main() {
// 	nums1 := []int{3, 2, 2, 3}

// 	log.Println(removeElement(nums1, 3), nums1)

// 	nums1 = []int{0, 1, 2, 2, 3, 0, 4, 2}

// 	log.Println(removeElement(nums1, 2), nums1)

// 	nums1 = []int{}

// 	log.Println(removeElement(nums1, 3), nums1)

// 	nums1 = []int{1, 2, 3, 4, 5}

// 	log.Println(removeElement(nums1, 10), nums1)

// 	nums1 = []int{1}

// 	log.Println(removeElement(nums1, 1), nums1)
// }
