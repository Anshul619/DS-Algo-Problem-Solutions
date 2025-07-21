package main

import "testing"

func TestHindex(t *testing.T) {
	cases := []struct {
		citations []int
		expected  int
	}{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{0, 0, 0}, 0},
		{[]int{10, 8, 5, 4, 3}, 4},
		{[]int{}, 0},
	}

	for _, c := range cases {
		result := hIndex(c.citations)
		if result != c.expected {
			t.Errorf("hIndex(%v) = %d; expected %d", c.citations, result, c.expected)
		}
	}
}
