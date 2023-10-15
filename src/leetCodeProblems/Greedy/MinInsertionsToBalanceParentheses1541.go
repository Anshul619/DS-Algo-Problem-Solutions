package main

/*
- LeetCode - https://leetcode.com/problems/optimal-partition-of-string/description/
- Time - O(n)
- Space - O(1)
*/
func minInsertions(s string) int {

	c_needed := 0
	o_needed := 0

	for _, v := range s {
		if string(v) == "(" {
			c_needed += 2

			if c_needed%2 == 1 {
				c_needed--
				o_needed++
			}
		} else {
			c_needed -= 1

			if c_needed == -1 {
				c_needed = 1
				o_needed++
			}
		}
	}

	return c_needed + o_needed
}

// func main() {

// 	//s := "abacaba"

// 	s := "ssssss"
// 	log.Println(partitionString(s))
// }
