package main

/*
- LeetCode - https://leetcode.com/problems/spiral-matrix/
- Time - O(m*n)
- Space - O(m*n)
*/
func spiralOrder(matrix [][]int) []int {
	dX := []int{0, 1, 0, -1}
	dY := []int{1, 0, -1, 0}

	direction := 0
	curRow := 0
	curCol := 0

	out := []int{}
	m := make([][]bool, len(matrix))

	for i := 0; i < len(matrix); i++ {
		m[i] = make([]bool, len(matrix[0]))
	}

	for i := 0; i < len(matrix)*len(matrix[0]); i++ {
		out = append(out, matrix[curRow][curCol])
		m[curRow][curCol] = true

		nextRow := curRow + dX[direction]
		nextCol := curCol + dY[direction]

		if nextRow >= 0 && nextRow < len(matrix) &&
			nextCol >= 0 && nextCol < len(matrix[0]) &&
			!m[nextRow][nextCol] {
			curRow = nextRow
			curCol = nextCol
		} else {
			direction = (direction + 1) % 4

			curRow += dX[direction]
			curCol += dY[direction]
		}
	}
	return out
}
