package main

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	for i, v := range tests {
		if coinChange(v.coins, v.amount) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
