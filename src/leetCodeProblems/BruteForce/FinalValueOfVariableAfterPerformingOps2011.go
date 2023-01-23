package main

/*
- LeetCode - https://leetcode.com/problems/final-value-of-variable-after-performing-operations/submissions/
- Time - O(n)
- Space - O(1)
*/

func finalValueAfterOperations(operations []string) int {

	X := 0

	for _, v := range operations {
		s := string(v)

		if s == "--X" || s == "X--" {
			X--
		} else if s == "++X" || s == "X++" {
			X++
		}
	}

	return X
}

// func main() {
// 	op := []string{"--X", "X++", "X++"}

// 	log.Println(finalValueAfterOperations(op))
// }
