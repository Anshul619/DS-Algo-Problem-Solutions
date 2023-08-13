package main

/*
- Leetcode - https://leetcode.com/problems/degree-of-an-array/description/
- Time - O(n)
- Space - O(n)
*/
import (
	"math"
)

type Num struct {
	Count      int
	StartIndex int
	EndIndex   int
}

func findShortestSubArray(nums []int) int {

	m := make(map[int]Num)
	degree, out := 0, math.MaxInt

	for i, v := range nums {
		if elem, ok := m[v]; ok {
			elem.Count++
			elem.EndIndex = i
			m[v] = elem

			if elem.Count > degree {
				degree = elem.Count
				out = i - elem.StartIndex
			} else if elem.Count == degree && i-elem.StartIndex < out {
				out = i - elem.StartIndex
			}

		} else {
			m[v] = Num{
				Count:      1,
				StartIndex: i,
				EndIndex:   i,
			}
		}
	}

	if out == math.MaxInt {

		if len(nums) > 0 {
			return 1
		}

		return 0
	}

	return out + 1
}

// func main() {
// 	log.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
// 	log.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
// 	log.Println(findShortestSubArray([]int{1}))
// }
