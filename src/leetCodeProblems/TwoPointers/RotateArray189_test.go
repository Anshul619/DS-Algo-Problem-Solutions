package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7}
		rotate(input, 3)
		if !reflect.DeepEqual(input, []int{5, 6, 7, 1, 2, 3, 4}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		input := []int{-1}
		rotate(input, 2)
		if !reflect.DeepEqual(input, []int{-1}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		input := []int{1, 2}
		rotate(input, 3)
		if !reflect.DeepEqual(input, []int{2, 1}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		input := []int{1, 2, 3}
		rotate(input, 4)
		if !reflect.DeepEqual(input, []int{3, 1, 2}) {
			t.Fail()
		}
	})

	t.Run("test5", func(t *testing.T) {
		input := []int{1, 2}
		rotate(input, 2)
		if !reflect.DeepEqual(input, []int{1, 2}) {
			t.Fail()
		}
	})

	t.Run("test6", func(t *testing.T) {
		input := []int{1, 2}
		rotate(input, 5)
		if !reflect.DeepEqual(input, []int{2, 1}) {
			t.Fail()
		}
	})
}
