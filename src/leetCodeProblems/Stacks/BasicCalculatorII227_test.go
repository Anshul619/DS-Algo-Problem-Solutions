package main

/*
- LeetCode - https://leetcode.com/problems/basic-calculator-ii/
*/

import (
	"reflect"
	"testing"
)

func TestBasicCal1(t *testing.T) {
	s := "3+2*2"
	expectedOut := 7

	if !reflect.DeepEqual(calculate(s), expectedOut) {
		t.Fatalf("TestBasicCal1 Case Failed")
	}
}

func TestBasicCal2(t *testing.T) {
	s := " 3+5 / 2 "
	expectedOut := 5

	if !reflect.DeepEqual(calculate(s), expectedOut) {
		t.Fatalf("TestBasicCal2 Case Failed")
	}
}

func TestBasicCal3(t *testing.T) {
	s := " 3/2 "
	expectedOut := 1

	if !reflect.DeepEqual(calculate(s), expectedOut) {
		t.Fatalf("TestBasicCal3 Case Failed")
	}
}

func TestBasicCal4(t *testing.T) {
	s := "42"
	expectedOut := 42

	if !reflect.DeepEqual(calculate(s), expectedOut) {
		t.Fatalf("TestBasicCal4 Case Failed")
	}
}
