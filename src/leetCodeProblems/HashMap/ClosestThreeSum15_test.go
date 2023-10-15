package main

import (
	"reflect"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		if !reflect.DeepEqual(threeSumClosest([]int{-1, 2, 1, -4}, 1), 2) {
			t.Fatalf("test1 failed")
		}
	})

	t.Run("test2", func(t *testing.T) {
		if !reflect.DeepEqual(threeSumClosest([]int{0, 0, 0}, 1), 0) {
			t.Fatalf("test2 failed")
		}
	})
	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2), -2) {
			t.Fatalf("test3 failed")
		}
	})
}
