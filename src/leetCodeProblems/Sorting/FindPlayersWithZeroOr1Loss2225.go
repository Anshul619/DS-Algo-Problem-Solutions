package main

/*
- LeetCode - https://leetcode.com/problems/find-players-with-zero-or-one-losses
- Time - O(nlogn)
- Space - O(n)
*/
import (
	"sort"
)

func findWinners(matches [][]int) [][]int {

	out := [][]int{}
	out = append(out, []int{}, []int{})

	losersCountMap := make(map[int]int)

	for _, v := range matches {

		if _, ok := losersCountMap[v[0]]; !ok {
			losersCountMap[v[0]] = 0
		}

		losersCountMap[v[1]]++
	}

	for k, v := range losersCountMap {

		if v == 0 {
			out[0] = append(out[0], k)
		}
		if v == 1 {
			out[1] = append(out[1], k)
		}
	}

	sort.Ints(out[0])
	sort.Ints(out[1])

	return out
}

// func main() {

// 	matches := [][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}

// 	//matches := [][]int{{2, 3}, {1, 3}, {5, 4}, {6, 4}}
// 	log.Println(findWinners(matches))
// }
