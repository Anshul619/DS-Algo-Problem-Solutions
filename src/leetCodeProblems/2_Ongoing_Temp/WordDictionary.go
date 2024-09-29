package main

type WordNode struct {
	children    [26]*WordNode
	isEndOfWord bool
}

type WordDictionary struct {
	root *WordNode
}

func Constructor() WordDictionary {
	w := WordDictionary{}
	w.root = new(WordNode)
	return w
}

func (this *WordDictionary) AddWord(word string) {
	w := this.root

	for _, v := range word {
		v -= 'a'

		if w.children[v] == nil {
			w.children[v] = new(WordNode)
		}

		w = w.children[v]
	}

	w.isEndOfWord = true
}

func searchUtil(word string, root *WordNode) bool {
	if root == nil {
		return false
	}

	w := root

	for i, v := range word {
		if v == '.' {
			for j := 0; j < 26; j++ {
				if w.children[j] != nil {
					if searchUtil(word[i+1:], w.children[j]) {
						return true
					}
				}
			}
			return false
		}

		v -= 'a'

		if w.children[v] == nil {
			return false
		}
		w = w.children[v]
	}
	return w != nil && w.isEndOfWord
}

func (this *WordDictionary) Search(word string) bool {
	return searchUtil(word, this.root)
}
