package main

import (
	"reflect"
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct {
		nums int
		out  int
	}{
		{4, 2},
		{8, 2},
		{1, 1},
		{3, 1},
		{0, 0},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(mySqrt(v.nums), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
