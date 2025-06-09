package main

import (
	"reflect"
	"testing"
)

/*
- LeetCode - https://leetcode.com/problems/game-of-life/
*/

func TestGameOfLife(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			input: [][]int{
				{1, 1},
				{1, 0},
			},
			expected: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			input: [][]int{
				{1},
			},
			expected: [][]int{
				{0},
			},
		},
	}

	for i, v := range tests {
		gameOfLife(v.input)
		if !reflect.DeepEqual(v.input, v.expected) {
			t.Errorf("failed %v", i)
		}
	}
}
