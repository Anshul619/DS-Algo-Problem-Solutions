package main

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path string
		out  string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/..", "/"},
		{"/a/../", "/"},
		{"/../../../../../a", "/a"},
		{"/a/./b/./c/./d/", "/a/b/c/d"},
		{"/a/../.././../../.", "/"},
		{"/a//b//c//////d", "/a/b/c/d"},
	}

	for i, v := range tests {
		if simplifyPath(v.path) != v.out {
			t.Errorf("failed %v", i)
		}
	}
}
