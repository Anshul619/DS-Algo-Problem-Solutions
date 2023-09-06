package main

/*
- Leetcode - https://leetcode.com/problems/merge-intervals/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastIndex := 0

	for i := 1; i < len(intervals); i++ {

		if intervals[lastIndex][1] >= intervals[i][0] {
			intervals[lastIndex][1] = max(intervals[i][1], intervals[lastIndex][1])
		} else {
			lastIndex++
			intervals[lastIndex] = intervals[i]
		}
	}
	return intervals[:lastIndex+1]
}

// func main() {
// 	log.Println(merge([][]int{{1, 3}, {15, 18}, {2, 6}, {8, 10}}))
// 	//log.Println(merge([][]int{{1, 4}, {4, 5}}))
// }
