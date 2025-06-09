package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/
- Space - O(m*n)
- Time - O(m*n)
*/

func winner(w int) string {
	switch w {
	case 1:
		return "A"
	case 2:
		return "B"
	default:
		return "None"
	}
}

func tictactoe(moves [][]int) string {
	board := make([][]int, 3)

	for i := range board {
		board[i] = make([]int, 3)
	}

	for i, v := range moves {
		if i%2 == 0 {
			board[v[0]][v[1]] = 1
		} else {
			board[v[0]][v[1]] = 2
		}
	}

	rowsCheck := checkRows(board)
	if rowsCheck != -1 {
		return winner(rowsCheck)
	}

	columnsCheck := checkColumns(board)
	log.Println(columnsCheck)
	if columnsCheck != -1 {
		return winner(columnsCheck)
	}

	diagonalsCheck := checkDiagonal(board)
	if diagonalsCheck != -1 {
		return winner(diagonalsCheck)
	}

	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}

func checkRows(g [][]int) int {
	for _, v := range g {
		if v[2] != 0 && v[0] == v[1] && v[1] == v[2] {
			return v[0]
		}
	}
	return -1
}
func checkColumns(g [][]int) int {
	for i := 0; i < 3; i++ {
		if g[0][i] != 0 && g[0][i] == g[1][i] && g[1][i] == g[2][i] {
			return g[0][i]
		}
	}
	return -1
}

func checkDiagonal(g [][]int) int {
	if g[0][0] != 0 && g[0][0] == g[1][1] && g[1][1] == g[2][2] {
		return g[0][0]
	}
	if g[2][0] != 0 && g[2][0] == g[1][1] && g[1][1] == g[0][2] {
		return g[2][0]
	}
	return -1
}

// func main() {
// 	// log.Println(tictactoe([][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}))
// 	// log.Println(tictactoe([][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}))
// 	// log.Println(tictactoe([][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}))
// 	log.Println(tictactoe([][]int{{2, 0}, {1, 1}, {0, 2}, {2, 1}, {1, 2}, {1, 0}, {0, 0}, {0, 1}}))

// }
