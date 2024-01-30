package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		if !isPalindrome(121) {
			t.Fail()
		}
	})

	t.Run("test2", func(t *testing.T) {
		if isPalindrome(-121) {
			t.Fail()
		}
	})

	t.Run("test3", func(t *testing.T) {
		if isPalindrome(10) {
			t.Fail()
		}
	})

	t.Run("test4", func(t *testing.T) {
		if !isPalindrome(0) {
			t.Fail()
		}
	})
}
