package main

/*
- LeetCode - https://leetcode.com/problems/game-of-life/
- Time - O(m*n)
- Space - O(1)
*/

func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}

	numRows := len(board)
	numCols := len(board[0])

	dx := []int{-1, 0, 1, 0, -1, 1, 1, -1}
	dy := []int{0, -1, 0, 1, 1, 1, -1, -1}

	for i, v := range board {
		for j, v1 := range v {
			neighLives := 0

			for k := 0; k < 8; k++ {
				if i+dx[k] > -1 && i+dx[k] < numRows &&
					j+dy[k] > -1 && j+dy[k] < numCols {
					if board[i+dx[k]][j+dy[k]] >= 1 {
						neighLives++
					}
				}
			}

			if v1 > 0 {
				board[i][j] += neighLives
			} else {
				board[i][j] -= neighLives
			}
		}
	}

	for i, v := range board {
		for j, v1 := range v {
			if v1 > 0 {

				v1 -= 1

				switch {
				case v1 < 2:
					board[i][j] = 0
				case v1 == 2 || v1 == 3:
					board[i][j] = 1
				default:
					board[i][j] = 0
				}
			} else {
				if v1 == -3 {
					board[i][j] = 1
				} else {
					board[i][j] = 0
				}
			}
		}
	}
}
