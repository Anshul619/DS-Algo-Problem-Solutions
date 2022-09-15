package main

/*
- LeetCode - https://leetcode.com/problems/continuous-subarray-sum
- Time Complexity - O(n)
- Space Complexity - O(min(n,k))
*/
import "log"

func checkSubarraySum(nums []int, k int) bool {

	m := map[int]int{
		0: 0,
	}

	currentSum := 0

	for i, v := range nums {

		currentSum += v

		if remainderIndex, ok := m[currentSum%k]; ok {
			if i > remainderIndex {
				return true
			}
		} else {
			m[currentSum%k] = i + 1
		}
	}
	return false
}

func main() {

	//nums := []int{23, 2, 4, 6, 7}
	//nums := []int{23, 2, 4, 6, 6}
	//nums := []int{5, 0, 0, 0}
	nums := []int{23, 2, 4, 6, 6}
	log.Println(checkSubarraySum(nums, 7))

}
