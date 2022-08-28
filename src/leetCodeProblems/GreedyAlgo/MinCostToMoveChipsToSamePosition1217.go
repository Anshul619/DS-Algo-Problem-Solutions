package main

/*
- LeetCode - https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/solution/
*/

import (
	"log"
)

func minCostToMoveChips(position []int) int {

	evenCount, oddCount := 0, 0

	for _, v := range position {
		if v%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	out := evenCount

	// The min count would have to be moved in 1 cost.
	if oddCount < evenCount {
		out = oddCount
	}

	return out
}

func main() {
	//input := []int{2, 2, 2, 3, 3}
	//input := []int{1, 2, 3}
	//input := []int{3, 3, 1, 2, 2}
	input := []int{6, 4, 7, 8, 2, 10, 2, 7, 9, 7}
	log.Println(minCostToMoveChips(input))
}
