package main

/*
- LeetCode - https://leetcode.com/problems/lru-cache/
- Time - O(1)
- Space - O(n)
*/

import "container/list"

type KeyValue struct {
	Key int
	Val int
}

type LRUCache struct {
	m        map[int]*list.Element
	l        *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		make(map[int]*list.Element),
		list.New(),
		capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
		this.l.MoveToFront(v)
		return v.Value.(KeyValue).Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.m[key]; ok {
		v.Value = KeyValue{key, value}
		this.l.MoveToFront(v)
		return
	}

	if this.l.Len() == this.capacity {
		e := this.l.Back()
		this.l.Remove(e)
		delete(this.m, e.Value.(KeyValue).Key)
	}
	e := this.l.PushFront(KeyValue{key, value})
	this.m[key] = e
}
