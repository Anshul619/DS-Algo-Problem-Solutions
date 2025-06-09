package main

/*
- Leetcode - https://leetcode.com/problems/rotate-image/description/
- Time - O(n/2)
- Space - O(1)
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	cycles := n / 2

	for i := 0; i < cycles; i++ {
		for j := i; j < n-i-1; j++ {

			t := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = t

		}
	}
}
