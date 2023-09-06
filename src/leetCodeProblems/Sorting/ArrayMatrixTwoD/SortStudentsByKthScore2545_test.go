package main

import (
	"reflect"
	"testing"
)

func TestSortTheStudents1(t *testing.T) {
	score := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}

	if !reflect.DeepEqual(sortTheStudents(score, 2), [][]int{{7, 5, 11, 2}, {10, 6, 9, 1}, {4, 8, 3, 15}}) {
		t.Fatalf("Failed case")
	}
}

func TestSortTheStudents2(t *testing.T) {
	score := [][]int{{3, 4}, {5, 6}}

	if !reflect.DeepEqual(sortTheStudents(score, 0), [][]int{{5, 6}, {3, 4}}) {
		t.Fatalf("Failed case")
	}
}
