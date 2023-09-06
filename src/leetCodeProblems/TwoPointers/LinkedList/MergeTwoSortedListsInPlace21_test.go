package main

import (
	"reflect"
	"testing"
)

func Test_MergeListInPlace1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)

	if !reflect.DeepEqual(nums1, []int{1, 2, 2, 3, 5, 6}) {
		t.Fatalf("Failed Test_MergeListInPlace1")
	}
}

func Test_MergeListInPlace2(t *testing.T) {
	nums1 := []int{1}
	merge(nums1, 1, []int{}, 0)

	if !reflect.DeepEqual(nums1, []int{1}) {
		t.Fatalf("Failed Test_MergeListInPlace2")
	}
}

func Test_MergeListInPlace3(t *testing.T) {
	nums1 := []int{0}
	merge(nums1, 0, []int{1}, 1)

	if !reflect.DeepEqual(nums1, []int{1}) {
		t.Fatalf("Failed Test_MergeListInPlace3")
	}
}

func Test_MergeListInPlace4(t *testing.T) {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	merge(nums1, 3, []int{1, 2, 3}, 3)

	if !reflect.DeepEqual(nums1, []int{1, 2, 3, 4, 5, 6}) {
		t.Fatalf("Failed Test_MergeListInPlace3")
	}
}
