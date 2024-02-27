package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for i, v := range tests {
		if maxProfit(v.prices) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
