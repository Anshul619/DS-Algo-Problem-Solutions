package main

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		out  int
	}{
		{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, 1},
	}

	for i, v := range tests {
		if numIslands(v.grid) != v.out {
			t.Errorf("Failed test %v", i)
		}
	}
}
