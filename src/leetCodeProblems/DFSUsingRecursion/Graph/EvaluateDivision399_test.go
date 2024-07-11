package main

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		out       []float64
	}{
		{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}, []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}},
		{[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0}, [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}}, []float64{3.75000, 0.40000, 5.00000, 0.20000}},
		{[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}, []float64{360.00000, 0.00833, 20.00000, 1.00000, -1.00000, -1.00000}},
		{[][]string{{"a", "b"}, {"e", "f"}, {"b", "e"}}, []float64{3.4, 1.4, 2.3}, [][]string{{"b", "a"}, {"a", "f"}, {"f", "f"}, {"e", "e"}, {"c", "c"}, {"a", "c"}, {"f", "e"}}, []float64{0.29412, 10.94800, 1.00000, 1.00000, -1.00000, -1.00000, 0.71429}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(calcEquation(v.equations, v.values, v.queries), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
