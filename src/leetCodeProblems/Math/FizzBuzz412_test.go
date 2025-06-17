package main

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {

	tests := []struct {
		n        int
		expected []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fizzBuzz(tt.n); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.expected)
			}
		})
	}
}
