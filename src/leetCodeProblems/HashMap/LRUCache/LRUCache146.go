package main

/*
- LeetCode - https://leetcode.com/problems/lru-cache/
- Time - O(1)
- Space - O(n)
*/

import "container/list"

type Node struct {
	key int
	val int
}

type LRUCache struct {
	l    *list.List
	m    map[int]*list.Element
	size int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size: capacity,
		l:    list.New(),
		m:    make(map[int]*list.Element),
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.m[key]; ok {
		this.l.MoveToBack(e)
		return e.Value.(*Node).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// ✅ If key exists → update + move
	if e, ok := this.m[key]; ok {
		node := e.Value.(*Node)
		node.val = value

		this.l.MoveToBack(e)
		return
	}

	// ✅ If full → evict LRU
	if this.l.Len() == this.size {
		f := this.l.Front()

		if f != nil {
			this.l.Remove(f)
			delete(this.m, f.Value.(*Node).key)
		}
	}

	// ✅ Insert new
	node := &Node{
		key: key,
		val: value,
	}

	this.m[key] = this.l.PushBack(node)
}
