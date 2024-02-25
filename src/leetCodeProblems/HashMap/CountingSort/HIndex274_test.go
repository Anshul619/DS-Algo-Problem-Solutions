package main

import "testing"

func TestHIndex(t *testing.T) {
	tests := []struct {
		citations []int
		expected  int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{0}, 0},
	}

	for i, v := range tests {
		if hIndex(v.citations) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
