package main

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		s        string
		numRows  int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for i, v := range tests {
		if convert(v.s, v.numRows) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
