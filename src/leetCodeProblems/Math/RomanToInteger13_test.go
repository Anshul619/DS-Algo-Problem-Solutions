package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, v := range tests {
		if romanToInt(v.s) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
