package main

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost     []int
		expected int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for i, v := range tests {
		if minCostClimbingStairs(v.cost) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
