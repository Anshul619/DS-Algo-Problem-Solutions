package main

import "testing"

func TestMaxProfit1(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for i, v := range tests {
		if maxProfit1(v.prices) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
