package main

/*
- LeetCode - https://leetcode.com/problems/implement-trie-prefix-tree/solutions/3577631/go-hash-table-92-faster/
- Time - O(n)
- Space - O(n)
*/

type TrieNode struct {
	Children [26]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	trie := new(Trie)
	trie.root = new(TrieNode)
	return *trie
}

func (this *Trie) Insert(word string) {

	node := this.root

	for _, v := range word {
		v -= 'a'
		if node.Children[v] == nil {
			node.Children[v] = new(TrieNode)
		}
		node = node.Children[v]
	}

	node.isWord = true
}

func (this Trie) find(word string) *TrieNode {
	node := this.root

	for _, v := range word {
		v -= 'a'
		if node.Children[v] == nil {
			return nil
		}
		node = node.Children[v]
	}

	return node
}

func (this *Trie) Search(word string) bool {
	node := this.find(word)
	return node != nil && node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) != nil
}
