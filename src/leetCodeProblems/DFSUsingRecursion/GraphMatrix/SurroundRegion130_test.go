package main

import (
	"reflect"
	"testing"
)

func TestSurroundRegion(t *testing.T) {
	tests := []struct {
		board [][]byte
		out   [][]byte
	}{
		{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		{[][]byte{{'X'}}, [][]byte{{'X'}}},
		{[][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}, [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}},
	}

	for i, v := range tests {
		solve(v.board)
		if !reflect.DeepEqual(v.board, v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
