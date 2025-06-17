package main

/**
 * LeetCode - https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
 * - Time - O(1)
 * - Space - O(1)
 */
func countOdds(low int, high int) int {

	out := (high - low) / 2

	if high%2 == 1 || low%2 == 1 {
		out++
	}

	return out
}

// func main() {
// 	log.Println(countOdds(3, 7))
// 	log.Println(countOdds(8, 10))
// }
