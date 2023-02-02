package main

import "testing"

func Test1_MinMaxPairSum1877(t *testing.T) {
	nums := []int{3, 5, 2, 3}

	if minPairSum(nums) != 7 {
		t.Failed()
	}
}

func Test2_MinMaxPairSum1877(t *testing.T) {
	nums := []int{3, 5, 4, 2, 4, 6}

	if minPairSum(nums) != 8 {
		t.Failed()
	}
}

func Test3_MinMaxPairSum1877(t *testing.T) {
	nums := []int{3, 2, 4, 1, 1, 5, 1, 3, 5, 1}

	if minPairSum(nums) != 6 {
		t.Failed()
	}
}
