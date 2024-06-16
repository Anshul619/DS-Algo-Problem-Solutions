package main

import "testing"

func TestTribacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{4, 4},
		{25, 1389537},
	}

	for i, v := range tests {
		if tribonacci(v.n) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
