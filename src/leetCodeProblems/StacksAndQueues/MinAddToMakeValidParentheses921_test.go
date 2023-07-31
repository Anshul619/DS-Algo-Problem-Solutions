package main

import (
	"reflect"
	"testing"
)

func Test_minAddToMakeValid1(t *testing.T) {
	if !reflect.DeepEqual(minAddToMakeValid("())"), 1) {
		t.Fatalf("Test_minAddToMakeValid1 failed")
	}
}

func Test_minAddToMakeValid2(t *testing.T) {
	if !reflect.DeepEqual(minAddToMakeValid("((("), 3) {
		t.Fatalf("Test_minAddToMakeValid2 failed")
	}
}
