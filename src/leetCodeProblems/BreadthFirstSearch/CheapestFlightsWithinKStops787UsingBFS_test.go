package main

/*
- LeetCode - https://leetcode.com/problems/asteroid-collision/
*/

import (
	"reflect"
	"testing"
)

func TestCheapestPriceBFS1(t *testing.T) {

	input := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	expectedOutput := 700

	if !reflect.DeepEqual(expectedOutput, findCheapestPrice(4, input, 0, 3, 1)) {
		t.Fatalf("TestAsteriod1 failing")
	}
}

func TestCheapestPriceBFS2(t *testing.T) {

	input := [][]int{{0, 1, 100}, {1, 2, 100}}
	expectedOutput := 200

	if !reflect.DeepEqual(expectedOutput, findCheapestPrice(3, input, 0, 2, 1)) {
		t.Fatalf("TestAsteriod1 failing")
	}
}

func TestCheapestPriceBFS3(t *testing.T) {

	input := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	expectedOutput := 500

	if !reflect.DeepEqual(expectedOutput, findCheapestPrice(3, input, 0, 2, 0)) {
		t.Fatal("Failing")
	}
}

func TestCheapestPriceBFS4(t *testing.T) {

	input := [][]int{{4, 1, 1}, {1, 2, 3}, {0, 3, 2}, {0, 4, 10}, {3, 1, 1}, {1, 4, 3}}
	expectedOutput := -1

	if !reflect.DeepEqual(expectedOutput, findCheapestPrice(5, input, 2, 1, 1)) {
		t.Fatal("Failing")
	}
}

func TestCheapestPriceBFS5(t *testing.T) {

	input := [][]int{{3, 4, 4}, {2, 5, 6}, {4, 7, 10}, {9, 6, 5}, {7, 4, 4}, {6, 2, 10}, {6, 8, 6}, {7, 9, 4}, {1, 5, 4}, {1, 0, 4}, {9, 7, 3}, {7, 0, 5}, {6, 5, 8}, {1, 7, 6}, {4, 0, 9}, {5, 9, 1}, {8, 7, 3}, {1, 2, 6}, {4, 1, 5}, {5, 2, 4}, {1, 9, 1}, {7, 8, 10}, {0, 4, 2}, {7, 2, 8}}
	expectedOutput := 14

	if !reflect.DeepEqual(expectedOutput, findCheapestPrice(10, input, 6, 0, 7)) {
		t.Fatal("Failing")
	}
}

func TestCheapestPriceBFS6(t *testing.T) {

	input := [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2}, {4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}}
	expectedOutput := 5

	if !reflect.DeepEqual(expectedOutput, findCheapestPrice(5, input, 0, 4, 1)) {
		t.Fatal("Failing")
	}
}
