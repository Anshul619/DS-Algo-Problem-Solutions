package main

/*
- LeetCode - https://leetcode.com/problems/asteroid-collision/
*/

import (
	"reflect"
	"testing"
)

func TestAsteriod1(t *testing.T) {

	input := []int{5, 10, -5}
	expectedOutput := []int{5, 10}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatalf("TestAsteriod1 failing")
	}
}

func TestAsteriod2(t *testing.T) {

	input := []int{8, -8}
	expectedOutput := []int{}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatalf("TestAsteriod1 failing")
	}
}

func TestAsteriod3(t *testing.T) {

	input := []int{10, 2, -5}
	expectedOutput := []int{10}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatal("Output", asteroidCollision(input))
	}
}

func TestAsteriod4(t *testing.T) {

	input := []int{3, -2, 4}
	expectedOutput := []int{3, 4}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatal("Output", asteroidCollision(input))
	}
}

func TestAsteriod5(t *testing.T) {

	input := []int{-2, -1, 1, 2}
	expectedOutput := []int{-2, -1, 1, 2}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatal("Output", asteroidCollision(input))
	}
}

func TestAsteriod6(t *testing.T) {

	input := []int{-2, -2, 1, -2}
	expectedOutput := []int{-2, -2, -2}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatal("Output", asteroidCollision(input))
	}
}

func TestAsteriod7(t *testing.T) {

	input := []int{-2, -2, 1, -1}
	expectedOutput := []int{-2, -2}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatal("Output", asteroidCollision(input))
	}
}

func TestAsteriod8(t *testing.T) {

	input := []int{1, -2, -2, -2}
	expectedOutput := []int{-2, -2, -2}

	if !reflect.DeepEqual(expectedOutput, asteroidCollision(input)) {
		t.Fatal("Output", asteroidCollision(input))
	}
}
