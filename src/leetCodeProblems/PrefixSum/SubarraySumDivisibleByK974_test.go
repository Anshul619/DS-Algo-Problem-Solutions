package main

import (
	"reflect"
	"testing"
)

/*
- LeetCode - https://leetcode.com/problems/continuous-subarray-sum
*/

func TestSubarraySumDivisibleByK9741(t *testing.T) {

	input := []int{4, 5, 0, -2, -3, 1}
	expectedOutput := 7

	if !reflect.DeepEqual(expectedOutput, subarraysDivByK(input, 5)) {
		t.Fatalf("TestSubarraySumDivisibleByK9741 failing")
	}
}

func TestSubarraySumDivisibleByK9742(t *testing.T) {

	input := []int{5}
	expectedOutput := 0

	if !reflect.DeepEqual(expectedOutput, subarraysDivByK(input, 9)) {
		t.Fatalf("TestSubarraySumDivisibleByK9742 failing")
	}
}

func TestSubarraySumDivisibleByK9743(t *testing.T) {

	input := []int{-1, 2, 9}
	expectedOutput := 2

	if !reflect.DeepEqual(expectedOutput, subarraysDivByK(input, 2)) {
		t.Fatalf("TestSubarraySumDivisibleByK9743 failing")
	}
}
func TestSubarraySumDivisibleByK9744(t *testing.T) {

	input := []int{2, -2, 2, -4}
	expectedOutput := 2

	if !reflect.DeepEqual(expectedOutput, subarraysDivByK(input, 6)) {
		t.Fatalf("TestSubarraySumDivisibleByK9744 failing")
	}
}
