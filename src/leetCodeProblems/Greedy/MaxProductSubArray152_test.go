package main

/*
- LeetCode - https://leetcode.com/problems/maximum-product-subarray
*/
import (
	"reflect"
	"testing"
)

func TestProductSubArray1(t *testing.T) {

	input := []int{2, 3, -2, 4}
	expectedOutput := 6

	if !reflect.DeepEqual(expectedOutput, maxProduct(input)) {
		t.Fatalf("TestProductSubArray1")
	}
}

func TestProductSubArray2(t *testing.T) {

	input := []int{-2, 0, -1}
	expectedOutput := 0

	if !reflect.DeepEqual(expectedOutput, maxProduct(input)) {
		t.Fatalf("TestProductSubArray2")
	}
}

func TestProductSubArray3(t *testing.T) {

	input := []int{0, 2}
	expectedOutput := 2

	if !reflect.DeepEqual(expectedOutput, maxProduct(input)) {
		t.Fatalf("TestProductSubArray3")
	}
}

func TestProductSubArray4(t *testing.T) {

	input := []int{3, -1, 4}
	expectedOutput := 4

	if !reflect.DeepEqual(expectedOutput, maxProduct(input)) {
		t.Fatalf("TestProductSubArray3")
	}
}

func TestProductSubArray5(t *testing.T) {

	input := []int{2, -5, -2, -4, 3}
	expectedOutput := 24

	if !reflect.DeepEqual(expectedOutput, maxProduct(input)) {
		t.Fatalf("TestProductSubArray3")
	}
}
