package main

import "testing"

func TestEqualFrequency(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"abcc", true},
		{"aazz", false},
		{"bac", true},
		{"aca", true},
	}

	for i, v := range tests {
		if equalFrequency(v.s) != v.expected {
			t.Errorf("failed test %v", i)
		}
	}
}
