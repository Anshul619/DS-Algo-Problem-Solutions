package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperature(t *testing.T) {
	tests := []struct {
		temperatures []int
		out          []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(dailyTemperatures(v.temperatures), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
