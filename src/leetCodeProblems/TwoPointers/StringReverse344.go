package main

/*
- Leetcode - https://leetcode.com/problems/reverse-string/description/
- Time - O(n)
- Space - O(1)
*/
func reverseString(s []byte) {
	first, last := 0, len(s)-1

	for first < last {
		temp := s[first]
		s[first] = s[last]
		s[last] = temp
		first++
		last--
	}
}

// func main() {
// 	input := []byte{'h', 'e', 'h', 'e', 'l', 'l', 'o'}
// 	reverseString(input)
// 	log.Println(string(input))
// }
