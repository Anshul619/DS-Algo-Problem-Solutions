package main

/*
- LeetCode - https://leetcode.com/problems/smallest-even-multiple/submissions/
*/

func smallestEvenMultiple(n int) int {

	if n <= 2 {
		return 2
	}

	if n%2 == 0 {
		return n
	}

	return n * 2
}

// func main() {
// 	//n := 5
// 	n := 6
// 	log.Println(smallestEvenMultiple(n))
// }
