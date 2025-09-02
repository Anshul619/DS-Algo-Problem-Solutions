package main

/*
- LeetCode - https://leetcode.com/problems/implement-trie-prefix-tree/solutions/3577631/go-hash-table-92-faster/
- Time - O(L) (where L = length of word)
- Space - O(N*L) (where N = number of words, L = length of word)
*/

type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}
type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	t := Trie{}
	t.root = new(TrieNode)
	return t
}

func (this *Trie) Insert(word string) {
	node := this.root

	for _, v := range word {
		v -= 'a'
		if node.children[v] == nil {
			node.children[v] = new(TrieNode)
		}
		node = node.children[v]
	}
	node.isWord = true
}

func (this *Trie) Search(word string) bool {
	node := this.root

	for _, v := range word {
		v -= 'a'
		if node.children[v] == nil {
			return false
		}
		node = node.children[v]
	}

	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root

	for _, v := range prefix {
		v -= 'a'
		if node.children[v] == nil {
			return false
		}
		node = node.children[v]
	}

	return true
}
