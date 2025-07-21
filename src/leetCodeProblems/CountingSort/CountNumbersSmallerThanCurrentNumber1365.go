package main

/*
	LeetCode - https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
*/

//"log"

func smallerNumbersThanCurrent(nums []int) []int {
	counter := make([]int, 101)

	// Frequency counting (bucket sort idea)
	for _, v := range nums {
		counter[v]++
	}

	// Prefix sum to get cumulative counts
	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	// Cumulative count of all numbers smaller than nums[i]
	for i := range nums {
		if nums[i] != 0 {
			nums[i] = counter[nums[i]-1]
		}
	}
	return nums
}

// func main() {

// 	input := []int{8, 1, 2, 2, 3}

// 	log.Println(smallerNumbersThanCurrent(input))
// }
