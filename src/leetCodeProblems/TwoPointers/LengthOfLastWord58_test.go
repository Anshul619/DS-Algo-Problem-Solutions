package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
		{"a ", 1},
	}

	for i, v := range tests {
		if lengthOfLastWord(v.s) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
