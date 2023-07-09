package main

import (
	"reflect"
	"testing"
)

func Test_LongestSubstringWithoutRepeat1(t *testing.T) {
	input := "abcabcbb"
	expectedOut := 3
	if !reflect.DeepEqual(expectedOut, lengthOfLongestSubstring(input)) {
		t.Fatalf("Test_LongestSubstringWithoutRepeat1 failing")
	}
}

func Test_LongestSubstringWithoutRepeat2(t *testing.T) {
	input := "bbbbb"
	expectedOut := 1
	if !reflect.DeepEqual(expectedOut, lengthOfLongestSubstring(input)) {
		t.Fatalf("Test_LongestSubstringWithoutRepeat2 failing")
	}
}

func Test_LongestSubstringWithoutRepeat3(t *testing.T) {
	input := "pwwkew"
	expectedOut := 3
	if !reflect.DeepEqual(expectedOut, lengthOfLongestSubstring(input)) {
		t.Fatalf("Test_LongestSubstringWithoutRepeat3 failing")
	}
}

func Test_LongestSubstringWithoutRepeat4(t *testing.T) {
	input := "abba"
	expectedOut := 2
	if !reflect.DeepEqual(expectedOut, lengthOfLongestSubstring(input)) {
		t.Fatalf("Test_LongestSubstringWithoutRepeat4 failing")
	}
}

func Test_LongestSubstringWithoutRepeat5(t *testing.T) {
	input := " "
	expectedOut := 1
	if !reflect.DeepEqual(expectedOut, lengthOfLongestSubstring(input)) {
		t.Fatalf("Test_LongestSubstringWithoutRepeat5 failing")
	}
}
