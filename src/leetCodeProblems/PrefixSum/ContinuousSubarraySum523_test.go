package main

import (
	"reflect"
	"testing"
)

/*
- LeetCode - https://leetcode.com/problems/continuous-subarray-sum
*/

func TestContinuousSubarraySum1(t *testing.T) {

	input := []int{23, 2, 4, 6, 7}
	expectedOutput := true

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 6)) {
		t.Fatalf("ContinuousSubarraySum1 failing")
	}
}

func TestContinuousSubarraySum2(t *testing.T) {

	input := []int{23, 2, 6, 4, 7}
	expectedOutput := true

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 6)) {
		t.Fatalf("ContinuousSubarraySum2 failing")
	}
}

func TestContinuousSubarraySum3(t *testing.T) {

	input := []int{23, 2, 6, 4, 7}
	expectedOutput := false

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 13)) {
		t.Fatalf("ContinuousSubarraySum3 failing")
	}
}

func TestContinuousSubarraySum4(t *testing.T) {

	input := []int{5, 0, 0, 0}
	expectedOutput := true

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 3)) {
		t.Fatalf("ContinuousSubarraySum4 failing")
	}
}

func TestContinuousSubarraySum5(t *testing.T) {

	input := []int{0, 0}
	expectedOutput := true

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 1)) {
		t.Fatalf("ContinuousSubarraySum5 failing")
	}
}

func TestContinuousSubarraySum6(t *testing.T) {

	input := []int{0}
	expectedOutput := false

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 1)) {
		t.Fatalf("ContinuousSubarraySum5 failing")
	}
}

func TestContinuousSubarraySum7(t *testing.T) {

	input := []int{23, 2, 4, 6, 6}
	expectedOutput := true

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 7)) {
		t.Fatalf("ContinuousSubarraySum5 failing")
	}
}

func TestContinuousSubarraySum8(t *testing.T) {

	input := []int{0}
	expectedOutput := false

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 1)) {
		t.Fatalf("ContinuousSubarraySum8 failing")
	}
}

func TestContinuousSubarraySum9(t *testing.T) {

	input := []int{1, 0}
	expectedOutput := false

	if !reflect.DeepEqual(expectedOutput, checkSubarraySum(input, 2)) {
		t.Fatalf("TestContinuousSubarraySum9 failing")
	}
}
