package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
	}

	for i, v := range tests {
		if climbStairs(v.n) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
