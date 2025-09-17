package main

/*
- LeetCode - https://leetcode.com/problems/valid-sudoku/description/
- Time - O(9*9)
- Space - O(9)
*/
func isValidSudoku(board [][]byte) bool {

	// Rows traversal
	for i := 0; i < 9; i++ {
		m := make([]bool, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if ok := m[board[i][j]-byte('0')-1]; ok {
					return false
				}
				m[board[i][j]-byte('0')-1] = true
			}
		}
	}

	// Columns traversal
	for j := 0; j < 9; j++ {
		m := make([]bool, 9)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if ok := m[board[i][j]-byte('0')-1]; ok {
					return false
				}
				m[board[i][j]-byte('0')-1] = true
			}
		}
	}

	// Sub-boxes traversal
	rowStart := 0
	rowEnd := 2

	// first row boxes, then next
	for boxRow := 0; boxRow < 3; boxRow++ {
		colStart := 0
		colEnd := 2

		for boxCol := 0; boxCol < 3; boxCol++ {
			m := make([]bool, 9)

			for i := rowStart; i <= rowEnd; i++ {
				for j := colStart; j <= colEnd; j++ {
					if board[i][j] != '.' {
						if ok := m[board[i][j]-byte('0')-1]; ok {
							return false
						}
						m[board[i][j]-byte('0')-1] = true
					}
				}
			}

			colStart += 3
			colEnd += 3
		}

		rowStart += 3
		rowEnd += 3
	}
	return true
}
