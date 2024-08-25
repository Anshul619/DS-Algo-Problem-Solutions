package main

import (
	"testing"
)

func TestImplementTrie(t *testing.T) {

	trie := Constructor()
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Errorf("failed1")
	}

	if trie.Search("app") {
		t.Errorf("failed2")
	}

	if !trie.StartsWith("app") {
		t.Errorf("failed3")
	}

	trie.Insert("app")

	if !trie.Search("app") {
		t.Errorf("failed4")
	}
}
