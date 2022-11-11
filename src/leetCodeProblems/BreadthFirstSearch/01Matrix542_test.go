package main

import (
	"reflect"
	"testing"
)

func Test_updateMatrix1(t *testing.T) {

	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	expectedOutput := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}

	if !reflect.DeepEqual(expectedOutput, updateMatrix(mat)) {
		t.Fatalf("Test_updateMatrix1 failing")
	}
}

func Test_updateMatrix2(t *testing.T) {

	mat := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	expectedOutput := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	if !reflect.DeepEqual(expectedOutput, updateMatrix(mat)) {
		t.Fatalf("Test_updateMatrix2 failing")
	}
}

func Test_updateMatrix3(t *testing.T) {

	mat := [][]int{{0, 1, 0, 1, 1}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}}
	expectedOutput := [][]int{{0, 1, 0, 1, 2}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 0}, {1, 0, 1, 1, 1}, {1, 0, 0, 0, 1}}

	if !reflect.DeepEqual(expectedOutput, updateMatrix(mat)) {
		t.Fatalf("Test_updateMatrix3 failing")
	}
}

func Test_updateMatrix4(t *testing.T) {

	mat := [][]int{{1, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 1, 1, 1, 1}}
	expectedOutput := [][]int{{2, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 2, 2, 1}, {1, 1, 1, 0, 0, 1, 2, 2, 1, 0}, {0, 1, 2, 1, 0, 1, 2, 3, 2, 1}, {0, 0, 1, 2, 1, 2, 1, 2, 1, 0}, {1, 1, 2, 3, 2, 1, 0, 1, 1, 1}, {0, 1, 2, 3, 2, 1, 1, 0, 0, 1}, {1, 2, 1, 2, 1, 0, 0, 1, 1, 2}, {0, 1, 0, 1, 1, 0, 1, 2, 2, 3}, {1, 2, 1, 0, 1, 0, 1, 2, 3, 4}}

	if !reflect.DeepEqual(expectedOutput, updateMatrix(mat)) {
		t.Fatalf("Test_updateMatrix4 failing")
	}
}
