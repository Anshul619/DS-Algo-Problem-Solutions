package main

/*
- Leetcode - https://leetcode.com/problems/time-needed-to-buy-tickets/description/
- Time - O(n)
- Space - O(1)
*/
func timeRequiredToBuy(tickets []int, k int) int {
	out := tickets[k]

	for i := 0; i < k; i++ {
		out += min(tickets[i], tickets[k])
	}

	for i := k + 1; i < len(tickets); i++ {
		out += min(tickets[i], tickets[k]-1)
	}
	return out
}
