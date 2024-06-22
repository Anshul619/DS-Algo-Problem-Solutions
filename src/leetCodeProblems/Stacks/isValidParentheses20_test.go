package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"]", false},
	}

	for i, v := range tests {
		if isValid(v.s) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
