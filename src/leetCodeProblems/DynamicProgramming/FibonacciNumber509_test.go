package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for i, v := range tests {
		if fib(v.n) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
