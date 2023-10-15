package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/transpose-matrix/
- Time - O(n*m)
- Space - O(m*n)
*/

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}

	out := make([][]int, len(matrix[0]))
	for i := range out {
		out[i] = make([]int, len(matrix))
	}

	for i, v := range matrix {
		for j, v1 := range v {
			out[j][i] = v1
		}
	}
	return out
}

func main() {
	log.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	log.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
