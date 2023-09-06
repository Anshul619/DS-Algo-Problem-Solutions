package main

import (
	"reflect"
	"testing"
)

func Test_IsValidParenthesis(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		if !reflect.DeepEqual(isValid("()[]{}"), true) {
			t.Fatalf("Test_IsValidParenthesis failed")
		}
	})

	t.Run("test2", func(t *testing.T) {
		if !reflect.DeepEqual(isValid("()"), true) {
			t.Fatalf("Test_IsValidParenthesis failed")
		}
	})
	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(isValid("(]"), false) {
			t.Fatalf("Test_IsValidParenthesis failed")
		}
	})
	t.Run("test4", func(t *testing.T) {
		if !reflect.DeepEqual(isValid("[()]{}{[()()]()}"), true) {
			t.Fatalf("Test_IsValidParenthesis failed")
		}
	})
	t.Run("test5", func(t *testing.T) {
		if !reflect.DeepEqual(isValid("[(])"), false) {
			t.Fatalf("Test_IsValidParenthesis failed")
		}
	})
	t.Run("test6", func(t *testing.T) {
		if !reflect.DeepEqual(isValid("["), false) {
			t.Fatalf("Test_IsValidParenthesis failed")
		}
	})
}
