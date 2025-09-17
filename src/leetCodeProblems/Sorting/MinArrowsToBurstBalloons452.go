package main

/*
- Leetcode - https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/description/
- Time - O(nlogn)
- Space - O(1)
*/
import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	out := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			out++
			end = points[i][1]
		}
	}
	return out
}
