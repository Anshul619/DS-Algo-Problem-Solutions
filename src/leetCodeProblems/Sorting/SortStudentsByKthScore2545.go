package main

/*
- LeetCode - https://leetcode.com/problems/sort-the-students-by-their-kth-score/description/
- Time - O(n∗m∗logn)
- Space - O(1)
*/
import (
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.SliceStable(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}

// func main() {

// 	score := [][]int{{3, 4}, {7, 5, 11, 2}, {4, 8, 3, 15}}

// 	log.Println(sortTheStudents(score, 2))
// }
