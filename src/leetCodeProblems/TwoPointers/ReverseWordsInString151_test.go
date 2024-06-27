package main

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s   string
		out string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(reverseWords(v.s), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
