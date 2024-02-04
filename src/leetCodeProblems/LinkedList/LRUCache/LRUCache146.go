package main

/*
- LeetCode - https://leetcode.com/problems/lru-cache/
- Time - O(1)
- Space - O(n)
*/

import "container/list"

type KeyValue struct {
	key int
	val int
}

type LRUCache struct {
	m        map[int]*list.Element
	q        *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{make(map[int]*list.Element), list.New(), capacity}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.q.MoveToFront(n)
		return n.Value.(KeyValue).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	kv := KeyValue{key, value}

	if n, ok := this.m[key]; ok {
		n.Value = kv
		this.q.MoveToFront(n)
		this.m[key] = n
	} else {
		this.m[key] = this.q.PushFront(kv)
	}

	if len(this.m) > this.capacity {
		delete(this.m, this.q.Back().Value.(KeyValue).key)
		this.q.Remove(this.q.Back())
	}
}
