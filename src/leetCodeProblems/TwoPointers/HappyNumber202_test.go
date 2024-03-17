package main

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{19, true},
		{10, true},
		{2, false},
	}

	for i, v := range tests {
		if isHappy(v.n) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
