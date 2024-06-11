package main

import "testing"

func TestKLargest(t *testing.T) {
	obj := Constructor(3, []int{4, 5, 8, 2})

	tests := []struct {
		num      int
		expected int
	}{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	for i, v := range tests {
		if obj.Add(v.num) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
