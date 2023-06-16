package main

/*
- LeetCode - https://leetcode.com/problems/game-of-life/
- Time - O(m*n)
- Space - O(1)
*/

func checkExistence(board [][]int, x int, y int) bool {

	if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) {
		return true
	}

	return false
}

func gameOfLife(board [][]int) {

	dx := []int{1, 1, 0, -1, -1, -1, 0, 1}
	dy := []int{0, 1, 1, 1, 0, -1, -1, -1}

	for i, v := range board {

		for j, v1 := range v {

			if v1 > 0 {
				for k := 0; k < 8; k++ {
					if checkExistence(board, i+dx[k], j+dy[k]) && board[i+dx[k]][j+dy[k]] > 0 {
						board[i][j]++
					}
				}
			} else {
				for k := 0; k < 8; k++ {
					if checkExistence(board, i+dx[k], j+dy[k]) && board[i+dx[k]][j+dy[k]] > 0 {
						board[i][j]--
					}
				}
			}
		}
	}

	for i, v := range board {

		for j, v1 := range v {

			if v1 > 0 {
				if v1 < 3 {
					board[i][j] = 0
				} else if v1 == 3 || v1 == 4 {
					board[i][j] = 1
				} else if v1 > 4 {
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

// func main() {
// 	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}

// 	gameOfLife(board)

// 	log.Println(board)
// }
