package main

/*
- LeetCode - https://leetcode.com/problems/lru-cache/
- Time - O(1)
- Space - O(n)
*/

type CacheNode struct {
	Prev *CacheNode
	Next *CacheNode
	val  int
	key  int
}

type LRUCache1 struct {
	m        map[int]*CacheNode
	head     *CacheNode
	tail     *CacheNode
	capacity int
}

func (this *LRUCache1) PrintList() {
	list := []int{}
	temp := this.head
	for temp != nil {
		list = append(list, temp.val)
		temp = temp.Next
	}
}

func (this *LRUCache1) remove(node *CacheNode) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if this.head == node {
		this.head = this.head.Next
	}

	if this.tail == node {
		this.tail = this.tail.Prev
	}
}

func (this *LRUCache1) enqueue(key, value int) {
	node := new(CacheNode)
	node.val = value
	node.key = key

	node.Next = this.head

	if this.head != nil {
		this.head.Prev = node
	}

	if this.tail == nil {
		this.tail = node
	}
	this.head = node
	this.m[key] = node
}

func (this *LRUCache1) dequeue() {
	if this.tail == nil {
		return
	}

	lru := this.tail
	this.tail = this.tail.Prev
	this.tail.Next = nil
	delete(this.m, lru.key)
}

func Constructor1(capacity int) LRUCache1 {
	return LRUCache1{make(map[int]*CacheNode), nil, nil, capacity}
}

func (this *LRUCache1) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.remove(n)
		this.enqueue(key, n.val)
		return n.val
	}
	return -1
}

func (this *LRUCache1) Put(key int, value int) {
	if n, ok := this.m[key]; ok {
		this.remove(n)
	}

	this.enqueue(key, value)

	if len(this.m) > this.capacity {
		this.dequeue()
	}
}
