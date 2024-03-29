package main

/*
- LeetCode - https://leetcode.com/problems/strictly-palindromic-number/
- Time - O(n)
- Space - O(1)
*/
import (
	"strconv"
)

func calculateBase(base int, n int) string {
	out := ""
	temp := n
	for temp != 0 {
		out += strconv.Itoa(temp % base)
		temp = temp / base
	}
	return out
}

func isPalindrome1(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func isStrictlyPalindromic(n int) bool {

	for i := 2; i <= n-2; i++ {
		base := calculateBase(i, n)
		if !isPalindrome1(base) {
			return false
		}

	}
	//log.Println(calculateBase(2, n))
	return true
}

// func main() {
// 	log.Println(isStrictlyPalindromic(9))
// }
