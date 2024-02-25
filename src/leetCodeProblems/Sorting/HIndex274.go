package main

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)

	max := 0

	for i := len(citations) - 1; i >= 0; i-- {
		papers := len(citations) - i

		if citations[i] >= papers {
			max = papers
		} else {
			break
		}
	}
	return max
}
