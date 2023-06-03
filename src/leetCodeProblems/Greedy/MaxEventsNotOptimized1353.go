package main

/*
- LeetCode - https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/description/
- TimeExceeded
*/
import (
	"sort"
)

func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool {
		//return events[i][0] < events[j][0]
		return events[i][1] < events[j][1]
	})

	//log.Println(events)
	attended := make(map[int]bool)
	cur := 0

	for _, v := range events {

		if cur == 0 {
			cur = v[0]
			attended[cur] = true
		} else {
			ele := v[0]

			for ele <= v[1] {
				if _, ok := attended[ele]; !ok {
					cur = ele
					attended[ele] = true
					break
				}
				ele++
			}
		}
	}

	//log.Println(events)
	return len(attended)
}

// func main() {
// 	log.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}}))
// 	log.Println(maxEvents([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}))
// 	log.Println(maxEvents([][]int{{1, 1}, {2, 2}, {1, 2}}))
// 	log.Println(maxEvents([][]int{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}}))
// 	log.Println(maxEvents([][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}}))
// 	log.Println(maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}}))
// }
