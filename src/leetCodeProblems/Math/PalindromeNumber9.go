package main

/*
- Leetcode - https://leetcode.com/problems/palindrome-number/
- Time - O(n)
- Space - O(1)
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := 0
	reminder := x

	for reminder > 0 {
		reverse = reverse*10 + (reminder % 10)
		reminder = reminder / 10
	}

	return reverse == x
}
