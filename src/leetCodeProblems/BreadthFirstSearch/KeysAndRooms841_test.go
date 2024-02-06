package main

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		rooms    [][]int
		expected bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}

	for _, v := range tests {
		if canVisitAllRooms(v.rooms) != v.expected {
			t.Fail()
		}
	}
}
