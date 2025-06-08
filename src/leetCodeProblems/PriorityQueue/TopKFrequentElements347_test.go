package main

import (
	"testing"
)

func TestTopKFrequent1(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected map[int]bool
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, map[int]bool{1: true, 2: true}},
		{[]int{1, 2, 3, 4, 5}, 3, map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}},
		// Add more test cases...
	}

	for _, tt := range tests {
		res := topKFrequent1(tt.nums, tt.k)
		for _, r := range res {
			if !tt.expected[r] {
				t.Errorf("Unexpected element %d in result %v", r, res)
			}
		}
	}
}
