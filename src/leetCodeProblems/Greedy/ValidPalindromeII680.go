package main

/*
- LeetCode - https://leetcode.com/problems/valid-palindrome-ii/description/
- Time - O(n)
- Space - O(1)
*/

func isRemainingPalindrome(s string, left, right int) bool {
	if left < 0 || right >= len(s) {
		return false
	}
	for left < right {

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isRemainingPalindrome(s, left+1, right) || isRemainingPalindrome(s, left, right-1)
		}

		left++
		right--
	}

	return true
}

// func main() {
// 	log.Println(validPalindrome("aba"))
// 	log.Println(validPalindrome("abca"))
// 	log.Println(validPalindrome("abc"))
// }
