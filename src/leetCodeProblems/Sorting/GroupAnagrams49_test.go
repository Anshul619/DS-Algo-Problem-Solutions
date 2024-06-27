package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		out  [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{[]string{"eat", "tea", "ate"}, []string{"tan", "nat"}, []string{"bat"}}},
		{[]string{""}, [][]string{[]string{""}}},
	}

	for i, v := range tests {
		if !reflect.DeepEqual(groupAnagrams(v.strs), v.out) {
			t.Errorf("failed %v", i)
		}
	}
}
