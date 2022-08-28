package main

/**
- LeetCode - https://leetcode.com/problems/richest-customer-wealth/
- TimeComplexity - O(n)
- SpaceComplexity - O(n)
*/
import (
	//"log"
	"math"
)

func maximumWealth(accounts [][]int) int {

	var output, current float64

	for _, value := range accounts {

		current = 0
		for _, value1 := range value {
			current += float64(value1)
		}

		output = math.Max(output, current)
	}

	return int(output)
}

// func main() {

// 	input := [][]int{{10, 2, 3}, {3, 2, 5}}

// 	log.Println(maximumWealth(input))
// }
