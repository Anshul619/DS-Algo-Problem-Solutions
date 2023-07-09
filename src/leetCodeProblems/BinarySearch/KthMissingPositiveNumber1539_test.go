package main

import (
	"reflect"
	"testing"
)

func TestMissingNumber1(t *testing.T) {
	if !reflect.DeepEqual(findKthPositive([]int{2, 3, 4, 7, 11}, 5), 9) {
		t.Fatalf("TestMissingNumber1 failed")
	}
}

func TestMissingNumber2(t *testing.T) {
	if !reflect.DeepEqual(findKthPositive([]int{1, 2, 3, 4}, 2), 6) {
		t.Fatalf("TestMissingNumber2 failed")
	}
}

func TestMissingNumber3(t *testing.T) {
	if !reflect.DeepEqual(findKthPositive([]int{2}, 1), 1) {
		t.Fatalf("TestMissingNumber3 failed")
	}
}

func TestMissingNumber4(t *testing.T) {
	if !reflect.DeepEqual(findKthPositive([]int{7, 13, 21, 25, 29, 32, 38, 45}, 4), 4) {
		t.Fatalf("TestMissingNumber4 failed")
	}
}

func TestMissingNumber5(t *testing.T) {
	if !reflect.DeepEqual(findKthPositive([]int{1, 13, 18}, 17), 20) {
		t.Fatalf("TestMissingNumber5 failed")
	}
}
