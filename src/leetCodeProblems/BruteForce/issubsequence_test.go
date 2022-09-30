package main

import (
	"reflect"
	"testing"
)

func TestSubSequence1(t *testing.T) {

	s := "ahb"
	t1 := "ahbgdc"

	if !reflect.DeepEqual(isSubsequence(s, t1), true) {
		t.Fatal("Fail")
	}
}

func TestSubSequence2(t *testing.T) {

	s := "abc"
	t1 := "ahbgdc"

	if !reflect.DeepEqual(isSubsequence(s, t1), true) {
		t.Fatal("Fail")
	}
}

func TestSubSequence3(t *testing.T) {

	s := "axc"
	t1 := "ahbgdc"

	if !reflect.DeepEqual(isSubsequence(s, t1), false) {
		t.Fatal("Fail")
	}
}
