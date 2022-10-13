package main

/*
- LeetCode - https://leetcode.com/problems/valid-sudoku/submissions/
*/
import (
	"log"
	"strconv"
)

func isValidNum(num byte) bool {
	byteToInt, _ := strconv.Atoi(string(num))

	if 0 > byteToInt || byteToInt > 9 {
		log.Println("false")
		return false
	}

	return true
}

func isValidSudoku(board [][]byte) bool {

	mapDigits := make(map[byte]bool)

	rowsLength := len(board)
	colsLength := len(board[0])

	// Check all rows
	for i := 0; i < rowsLength; i++ {

		mapDigits = make(map[byte]bool)

		for j := 0; j < colsLength; j++ {

			if board[i][j] != '.' {
				if ok := mapDigits[board[i][j]]; ok || !isValidNum(board[i][j]) {
					return false
				}

				mapDigits[board[i][j]] = true
			}
		}
	}

	// Check all columns
	for i := 0; i < colsLength; i++ {

		mapDigits = make(map[byte]bool)

		for j := 0; j < rowsLength; j++ {

			if board[j][i] != '.' {
				//log.Println(i, j, board[i][j])
				if ok := mapDigits[board[j][i]]; ok || !isValidNum(board[j][i]) {
					return false
				}

				mapDigits[board[j][i]] = true
			}
		}
	}

	// Check subboxes 3 x 3
	for i := 0; i < rowsLength; i = i + 3 {

		for j := 0; j < colsLength; j = j + 3 {

			mapDigits = make(map[byte]bool)

			for k := 0; k < 3; k++ {

				for l := 0; l < 3; l++ {

					elem := board[i+k][j+l]

					if elem != '.' {
						if ok := mapDigits[elem]; ok || !isValidNum(elem) {
							return false
						}

						mapDigits[elem] = true
					}
				}
			}
		}
	}

	return true
}

func main() {

	//board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	board := [][]byte{{'.', '.', '4', '.', '.', '.', '6', '3', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'5', '.', '.', '.', '.', '.', '.', '9', '.'}, {'.', '.', '.', '5', '6', '.', '.', '.', '.'}, {'4', '.', '3', '.', '.', '.', '.', '.', '1'}, {'.', '.', '.', '7', '.', '.', '.', '.', '.'}, {'.', '.', '.', '5', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '.', '.', '.', '.'}}

	log.Println(isValidSudoku(board))
}
