package main

/*
- Leetcode - https://leetcode.com/problems/surrounded-regions/description/?envType=study-plan-v2&envId=top-interview-150
- Time - O(n * m)
- Space - O(n * m)
*/

func floodFill1(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}

	if board[x][y] != '-' {
		return
	}
	board[x][y] = 'O'

	// flood fill in all directions
	floodFill1(board, x+1, y)
	floodFill1(board, x, y+1)
	floodFill1(board, x-1, y)
	floodFill1(board, x, y-1)
}

func solve(board [][]byte) {
	for x, v := range board {
		for y, v1 := range v {
			if v1 == 'O' {
				board[x][y] = '-'
			}
		}
	}

	for x, v := range board {
		for y := range v {
			if x == 0 || y == 0 || x == len(board)-1 || y == len(v)-1 { // edges
				floodFill1(board, x, y)
			}
		}
	}

	for x, v := range board {
		for y := range v {
			if board[x][y] == '-' {
				board[x][y] = 'X'
			}
		}
	}
}
