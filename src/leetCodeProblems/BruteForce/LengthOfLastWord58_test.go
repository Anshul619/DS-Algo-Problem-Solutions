package main

import (
	"reflect"
	"testing"
)

func Test_LengthOfLastWord(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		if !reflect.DeepEqual(lengthOfLastWord("Hello World"), 5) {
			t.Fatalf("test1 failed")
		}
	})
	t.Run("test2", func(t *testing.T) {
		if !reflect.DeepEqual(lengthOfLastWord("   fly me   to   the moon  "), 4) {
			t.Fatalf("test2 failed")
		}
	})
	t.Run("test3", func(t *testing.T) {
		if !reflect.DeepEqual(lengthOfLastWord("luffy is still joyboy"), 6) {
			t.Fatalf("test3 failed")
		}
	})
	t.Run("test4", func(t *testing.T) {
		if !reflect.DeepEqual(lengthOfLastWord("Today is a nice day"), 3) {
			t.Fatalf("test4 failed")
		}
	})
}
