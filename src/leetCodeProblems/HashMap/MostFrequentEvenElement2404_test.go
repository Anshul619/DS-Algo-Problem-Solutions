package main

import "testing"

func TestMostFrequentEven(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{0, 1, 2, 2, 4, 4, 1}, 2},
		{[]int{4, 4, 4, 9, 2, 4}, 4},
		{[]int{29, 47, 21, 41, 13, 37, 25, 7}, -1},
	}

	for i, v := range tests {
		if mostFrequentEven(v.nums) != v.expected {
			t.Errorf("fail %v", i)
		}
	}
}
