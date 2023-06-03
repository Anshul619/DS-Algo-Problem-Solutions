package main

/*
- LeetCode - https://leetcode.com/problems/implement-trie-prefix-tree/solutions/3577631/go-hash-table-92-faster/
*/
import "log"

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
		i := v - 'a'
		if node.Children[i] == nil {
			node.Children[i] = new(TrieNode)
		}
		node = node.Children[i]
	}

	node.isWord = true
}

func (this Trie) find(word string) *TrieNode {
	node := this.root

	for _, v := range word {
		i := v - 'a'
		if node.Children[i] == nil {
			return nil
		}
		node = node.Children[i]
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

func main() {
	trie := Constructor()
	trie.Insert("apple")
	log.Println(trie.Search("apple"))

	log.Println(trie.Search("app"))
	log.Println(trie.StartsWith("app"))
	trie.Insert("app")
	log.Println(trie.Search("app"))
}
