package main

/*
- Leetcode - https://leetcode.com/problems/set-matrix-zeroes/
- Time - O(m*n)
- Space - O(1)
*/
func setZerosRow(matrix [][]int, row, startCol int) {
	for j := startCol; j < len(matrix[0]); j++ {
		matrix[row][j] = 0
	}
}

func setZerosCol(matrix [][]int, col, startRow int) {
	for i := startRow; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}

func setZeroes(matrix [][]int) {

	firstRowZero := false

	// Traverse the first row & set firstRowZero if its supposed to be zero
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	firstColZero := false

	// Traverse the second row & set firstColZero if its supposed to be zero
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// Traverse all rows & columns ( except first and last ) & mark first cell/row ZERO based on the condition
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Traverse the first column and set its rows ( not first ROW ) to ZERO based on condition
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			setZerosRow(matrix, i, 1)
		}
	}

	// Traverse the first row and set its columns ( not first column ) to ZERO based on condition
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			setZerosCol(matrix, j, 1)
		}
	}

	if firstRowZero {
		setZerosRow(matrix, 0, 0)
	}

	if firstColZero {
		setZerosCol(matrix, 0, 0)
	}
}
