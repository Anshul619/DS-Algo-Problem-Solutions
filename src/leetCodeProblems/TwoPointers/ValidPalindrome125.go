package main

import (
	"strings"
	"unicode"
)

/*
- LeetCode - https://leetcode.com/problems/valid-palindrome/
- Time - O(n)
- Space - O(1)
*/

func isAlphanumeric(b byte) bool {
	if unicode.IsLetter(rune(b)) || unicode.IsNumber(rune(b)) {
		return true
	}

	return false
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {

		if !isAlphanumeric(s[left]) {
			left++
			continue
		}

		if !isAlphanumeric(s[right]) {
			right--
			continue
		}

		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}

// Alternative solution
func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	f, l := 0, len(s)-1

	for f <= l {
		switch {

		case !unicode.IsLetter(rune(s[f])) && !unicode.IsDigit(rune(s[f])):
			f++
		case !unicode.IsLetter(rune(s[l])) && !unicode.IsDigit(rune(s[l])):
			l--
		default:
			if s[f] != s[l] {
				return false
			} else {
				f++
				l--
			}
		}
	}

	return true
}

// func main() {
// 	//plants := []int{2, 2, 3, 3}

// 	//log.Println(minimumRefill(plants, 5, 5))
// 	//log.Println(minimumRefill(plants, 3, 4))

// 	//plants := []int{5}
// 	log.Println(isPalindrome("A man, a plan, a canal: Panama"))
// 	log.Println(isPalindrome("race a car"))
// }
