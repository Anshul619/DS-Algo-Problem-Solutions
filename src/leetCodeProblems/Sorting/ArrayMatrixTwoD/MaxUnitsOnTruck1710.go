package main

import (
	"log"
	"sort"
)

/*
- LeetCode - https://leetcode.com/problems/maximum-units-on-a-truck/
- Time - O(nlogn)
- Space - O(1)
*/

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.SliceStable(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	out := 0

	for _, v := range boxTypes {
		if v[0] <= truckSize {
			out += v[0] * v[1]
			truckSize -= v[0]
		} else {
			out += truckSize * v[1]
			truckSize = 0
			break
		}
	}
	return out
}

func main() {
	log.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	log.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
}
