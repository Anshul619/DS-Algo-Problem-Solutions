package main

/*
- Leetcode - https://leetcode.com/problems/missing-number/
- Time - O(n)
- Space - O(1)
*/
func missingNumber(nums []int) int {
	currentSum := 0

	for _, v := range nums {
		currentSum += v
	}

	expectedSum := len(nums) * (len(nums) + 1) / 2
	return expectedSum - currentSum
}

// func main() {
// 	log.Println(missingNumber([]int{3, 0, 1}))
// 	log.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
// }
