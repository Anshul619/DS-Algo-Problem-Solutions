package main

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		input    []string
		k        int
		expected []string
	}{
		{[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(topKFrequent(v.input, v.k), v.expected) {
			t.Errorf("failed %v", i)
		}
	}
}
