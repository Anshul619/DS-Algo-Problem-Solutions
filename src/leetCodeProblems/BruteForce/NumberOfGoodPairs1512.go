package main

/*
- LeetCode - https://leetcode.com/problems/running-sum-of-1d-array/
*/
//import "log"

func numIdenticalPairs(nums []int) int {

	output := 0
	m := make(map[int]int)

	for _, value := range nums {
		if _, ok := m[value]; ok {
			output += m[value]
		}
		m[value]++
	}

	return output
}

func numIdenticalPairsBruteForce(nums []int) int {

	output := 0

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i] == nums[j] {
				output++
			}
		}
	}

	return output
}

// func main() {

// 	input := []int{1, 2, 3, 1, 1, 3}

// 	log.Println(numIdenticalPairs(input))

// }
