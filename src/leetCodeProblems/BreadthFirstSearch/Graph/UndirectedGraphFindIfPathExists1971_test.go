package main

import "testing"

func TestValidPath(t *testing.T) {
	tests := []struct {
		n           int
		edges       [][]int
		source      int
		destination int
		expected    bool
	}{
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2, true},
		{6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5, false},
	}

	for i, v := range tests {
		if validPath(v.n, v.edges, v.source, v.destination) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
