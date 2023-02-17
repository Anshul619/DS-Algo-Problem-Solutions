package main

/*
- LeetCode - https://leetcode.com/problems/count-of-matches-in-tournament/solutions/
- Time - O(n)
- Space - O(1)
*/

func numberOfMatches(n int) int {

	carryForward := n
	ans := 0

	for carryForward != 1 {
		if carryForward%2 == 0 {
			ans += carryForward / 2
			carryForward = carryForward / 2
		} else {
			ans += (carryForward - 1) / 2
			carryForward = 1 + (carryForward-1)/2
		}
	}

	return ans
}

// func main() {
// 	log.Println(numberOfMatches(7))
// 	log.Println(numberOfMatches(14))
// }
