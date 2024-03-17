package main

/*
- Leetcode - https://leetcode.com/problems/happy-number/
- Time - O(nlogn)
- Space - O(1)
*/

func calculateSquaresSum(n int) int {
	sum := 0

	for n != 0 {
		r := n % 10
		sum += r * r

		n = n / 10
	}
	return sum
}
func isHappy(n int) bool {

	slow, fast := n, n

	for {
		slow = calculateSquaresSum(slow)
		fast = calculateSquaresSum(calculateSquaresSum(fast))

		if slow == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}
