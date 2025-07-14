package main

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		heystack string
		needle   string
		out      int
	}{
		{"sadbutsad", "sad", 0},
		{"hello", "ll", 2},
		{"aaa", "aaaa", -1},
		{"mississippi", "issipi", -1},
	}

	for _, v := range tests {
		if strStr(v.heystack, v.needle) != v.out {
			t.Errorf("failed")
		}
	}
}
