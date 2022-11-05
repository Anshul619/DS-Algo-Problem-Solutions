package main

/**
- LeetCode - https://leetcode.com/problems/reverse-vowels-of-a-string/submissions/
- Time Complexity - O(n)
- Space Complexity - O(n)
**/
import "log"

func isVowel(s string) bool {

	if s == "a" || s == "e" || s == "i" || s == "o" || s == "u" ||
		s == "A" || s == "E" || s == "I" || s == "O" || s == "U" {
		return true
	}

	return false
}

func reverseVowels(s string) string {

	first := 0
	last := len(s) - 1
	output := []rune(s)

	for first < last {

		if !isVowel(string(s[first])) {
			first++
		}

		if !isVowel(string(s[last])) {
			last--
		}

		if isVowel(string(s[first])) && isVowel(string(s[last])) {
			output[first] = rune(s[last])
			output[last] = rune(s[first])
			first++
			last--
		}
	}

	return string(output)
}

func main() {
	//s := "hello"
	s := "leetcode"
	log.Println(reverseVowels(s))
}
