package main

/*
- LeetCode - https://leetcode.com/problems/asteroid-collision/
*/

import (
	"reflect"
	"testing"
)

func TestCheapestPriceDFS1(t *testing.T) {

	input := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	expectedOutput := 700

	if !reflect.DeepEqual(expectedOutput, findCheapestPriceDFS(4, input, 0, 3, 1)) {
		t.Fatalf("TestAsteriod1 failing")
	}
}

func TestCheapestPriceDFS2(t *testing.T) {

	input := [][]int{{0, 1, 100}, {1, 2, 100}}
	expectedOutput := 200

	if !reflect.DeepEqual(expectedOutput, findCheapestPriceDFS(3, input, 0, 2, 1)) {
		t.Fatalf("TestAsteriod1 failing")
	}
}

func TestCheapestPriceDFS3(t *testing.T) {

	input := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	expectedOutput := 500

	if !reflect.DeepEqual(expectedOutput, findCheapestPriceDFS(3, input, 0, 2, 0)) {
		t.Fatal("Failing")
	}
}
