package main

/*
- LeetCode - https://leetcode.com/problems/gas-station/description/
- Space - O(n)
- Time - O(1)
*/
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost := 0, 0

	for _, v := range gas {
		totalGas += v
	}

	for _, v := range cost {
		totalCost += v
	}

	if totalCost > totalGas {
		return -1
	}

	start, sum := 0, 0

	for i, v := range cost {
		sum += gas[i] - v

		if sum < 0 {
			start = i + 1
			sum = 0
		}
	}
	return start
}
