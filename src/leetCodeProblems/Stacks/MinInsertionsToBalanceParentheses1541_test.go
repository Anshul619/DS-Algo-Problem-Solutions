package main

import (
	"reflect"
	"testing"
)

func Test_minInsertions1(t *testing.T) {
	if !reflect.DeepEqual(minInsertions("(()))"), 1) {
		t.Fatalf("Test_minInsertions1 fail")
	}
}

func Test_minInsertions2(t *testing.T) {
	if !reflect.DeepEqual(minInsertions("())"), 0) {
		t.Fatalf("Test_minInsertions2 fail")
	}
}

func Test_minInsertions3(t *testing.T) {
	if !reflect.DeepEqual(minInsertions("))())("), 3) {
		t.Fatalf("Test_minInsertions3 fail")
	}
}

func Test_minInsertions4(t *testing.T) {
	t.Run("test41", func(t *testing.T) {
		if !reflect.DeepEqual(minInsertions("(()))(()))()())))"), 4) {
			t.Fatalf("Test_minInsertions4 fail")
		}
	})

	t.Run("test42", func(t *testing.T) {
		if !reflect.DeepEqual(minInsertions("(((()(()((())))(((()())))()())))(((()(()()((()()))"), 31) {
			t.Fatalf("Test_minInsertions4 fail")
		}
	})

}
