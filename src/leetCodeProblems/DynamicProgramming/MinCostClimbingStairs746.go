package main

/*
- Leetcode - https://leetcode.com/problems/min-cost-climbing-stairs/description/
- Time - O(n)
- Space - O(1)
*/

func min1(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func minCostClimbingStairs(cost []int) int {

	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min1(cost[i+1], cost[i+2])
	}
	return min1(cost[0], cost[1])
}

// func main() {
// 	log.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
// 	log.Println(minCostClimbingStairs([]int{10, 15, 20}))
// }
