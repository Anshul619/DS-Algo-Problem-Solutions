package main

import "testing"

func TestEvaluateRPN(t *testing.T) {
	tests := []struct {
		tokens []string
		out    int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for i, v := range tests {
		if evalRPN(v.tokens) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
