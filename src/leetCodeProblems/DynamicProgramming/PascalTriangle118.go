package main

/*
- LeetCode - https://leetcode.com/problems/pascals-triangle/
- Time - O(n^2)
- Space - O(n^2)
*/

func generate(numRows int) [][]int {
	out := [][]int{}

	for i := 1; i <= numRows; i++ {
		temp := make([]int, i)
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				temp[j-1] = 1
			} else {
				temp[j-1] = out[i-2][j-2] + out[i-2][j-1]
			}
		}
		out = append(out, temp)
	}
	return out
}

// func main() {
// 	//log.Println(minOperations(3))
// 	log.Println(generate(5))
// 	log.Println(generate(1))
// 	//log.Println(minOperations(31))
// }
