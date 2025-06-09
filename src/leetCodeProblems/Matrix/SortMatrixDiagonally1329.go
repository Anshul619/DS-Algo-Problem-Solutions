package main

/*
- LeetCode - https://leetcode.com/problems/sort-the-matrix-diagonally/submissions/
- TimeComplexity - O(nlogn)
- SpaceComplexity - O(1)
*/
import (
	"log"
	"sort"
)

func diagonalSort(mat [][]int) [][]int {

	rowsLength := len(mat)
	colsLength := len(mat[0])
	tempArray := []int{}

	for i := 0; i < rowsLength; i++ {

		j := 0

		for {

			if i+j < rowsLength && j < colsLength {
				tempArray = append(tempArray, mat[i+j][j])
				j++
			} else {
				break
			}

		}

		sort.Ints(tempArray)

		j = 0
		for {

			if i+j < rowsLength && j < colsLength {
				mat[i+j][j] = tempArray[j]
				j++
			} else {
				break
			}

		}

		tempArray = []int{}
	}

	for i := 1; i < colsLength; i++ {

		j := 0

		for {

			if j < rowsLength && i+j < colsLength {
				log.Println(mat[j][i+j])
				tempArray = append(tempArray, mat[j][i+j])
				j++
			} else {
				break
			}

		}

		sort.Ints(tempArray)

		j = 0

		for {

			if j < rowsLength && i+j < colsLength {
				mat[j][i+j] = tempArray[j]
				j++
			} else {
				break
			}

		}

		tempArray = []int{}
	}

	return mat
}

// func main() {
// 	mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
// 	log.Println(diagonalSort(mat))
// }
