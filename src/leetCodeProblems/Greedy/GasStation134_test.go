package main

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
	}

	for i, v := range tests {
		if canCompleteCircuit(v.gas, v.cost) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
