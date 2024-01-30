package main

/*
- Leetcode - https://leetcode.com/problems/merge-sorted-array
*/
import (
	"reflect"
	"testing"
)

func TestMergeSortedArray88(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		num1 := []int{1, 2, 3, 0, 0, 0}
		MergeArray(num1, 3, []int{2, 5, 6}, 3)
		if !reflect.DeepEqual(num1, []int{1, 2, 2, 3, 5, 6}) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		num1 := []int{1}
		MergeArray(num1, 1, []int{}, 0)
		if !reflect.DeepEqual(num1, []int{1}) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		num1 := []int{0}
		MergeArray(num1, 0, []int{1}, 1)
		if !reflect.DeepEqual(num1, []int{1}) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		nums1 := []int{4, 5, 6, 0, 0, 0}
		MergeArray(nums1, 3, []int{1, 2, 3}, 3)

		if !reflect.DeepEqual(nums1, []int{1, 2, 3, 4, 5, 6}) {
			t.Fail()
		}
	})
}
