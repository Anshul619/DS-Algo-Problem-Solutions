package main

/*
- LeetCode - https://leetcode.com/problems/game-of-life/
- Time - O(m*n)
- Space - O(m*n)
*/
func gameOfLife1(board [][]int) {
	if len(board) == 0 {
		return
	}

	numRows := len(board)
	numCols := len(board[0])
	out := make([][]int, numRows)

	for i := range out {
		out[i] = make([]int, numCols)
	}

	dx := []int{-1, 0, 1, 0, -1, 1, 1, -1}
	dy := []int{0, -1, 0, 1, 1, 1, -1, -1}

	for i, v := range board {
		for j, v1 := range v {

			neighLives := 0

			for k := 0; k < 8; k++ {
				if i+dx[k] > -1 && i+dx[k] < numRows &&
					j+dy[k] > -1 && j+dy[k] < numCols {
					if board[i+dx[k]][j+dy[k]] == 1 {
						neighLives++
					}
				}
			}

			switch v1 {
			case 1:
				switch {
				case neighLives < 2:
					out[i][j] = 0
				case neighLives == 2 || neighLives == 3:
					out[i][j] = 1
				default:
					out[i][j] = 0
				}
			default:
				if neighLives == 3 {
					out[i][j] = 1
				}
			}
		}
	}

	for i, v := range board {
		for j := range v {
			board[i][j] = out[i][j]
		}
	}
}
