package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		s   string
		out int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
		{"(5+(4+5+2)-3)+(6+8)", 27},
		{"2147483647", 2147483647},
		{"1-(     -2)", 3},
		{"-2+ 1", -1},
		{"1-11", -10},
	}

	for i, v := range tests {
		if calculate1(v.s) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
