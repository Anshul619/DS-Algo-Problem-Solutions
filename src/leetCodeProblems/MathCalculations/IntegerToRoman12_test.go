package main

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{3, "III"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i, v := range tests {
		if intToRoman(v.num) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
