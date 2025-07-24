package main

/*
- Leetcode - https://leetcode.com/problems/summary-ranges/description
*/
import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {

	tests := []struct {
		nums     []int
		expected []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}

	for _, v := range tests {
		if !reflect.DeepEqual(summaryRanges(v.nums), v.expected) {
			t.Fail()
		}
	}
}
