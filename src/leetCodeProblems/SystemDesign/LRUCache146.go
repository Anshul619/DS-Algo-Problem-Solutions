package main

/*
- LeetCode - https://leetcode.com/problems/lru-cache/
*/
// import (
// 	"log"
// )

type Queue []int

func (s *Queue) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Queue) Add(num int) {
	*s = append(*s, num)
}
func (s *Queue) Len() int {
	return len(*s)
}
func (s *Queue) Remove() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		elem := (*s)[0]
		*s = (*s)[1:]
		return elem, true
	}
}

func (s *Queue) RemoveKey(num int) {

	keyToBeRemoved := -1

	for i := range *s {
		if (*s)[i] == num {
			keyToBeRemoved = i
		}
	}

	if keyToBeRemoved == -1 {
		//Nothing
	} else if keyToBeRemoved == 0 {
		*s = (*s)[1:]
	} else if keyToBeRemoved == len(*s)-1 {
		*s = (*s)[:len(*s)-1]
	} else {
		*s = append((*s)[:keyToBeRemoved], (*s)[keyToBeRemoved+1:]...)
	}
}

type LRUCache struct {
	lruMap   map[int]int
	lruQueue *Queue
	size     int
}

func Constructor(capacity int) LRUCache {
	cache := new(LRUCache)
	cache.size = capacity
	cache.lruMap = make(map[int]int)
	cache.lruQueue = new(Queue)
	return *cache
}

func (this *LRUCache) addToQueue(key int) {

	if _, ok := this.lruMap[key]; ok {
		this.lruQueue.RemoveKey(key)
	}

	this.lruQueue.Add(key)
}

func (this *LRUCache) Get(key int) int {

	if v, ok := this.lruMap[key]; ok {
		this.addToQueue(key)
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.lruMap[key]; ok {
		this.addToQueue(key)
		this.lruMap[key] = value
		return
	}

	if this.size == len(this.lruMap) {
		remove, _ := this.lruQueue.Remove()
		delete(this.lruMap, remove)
	}

	this.addToQueue(key)
	this.lruMap[key] = value
}

func main() {
	// obj := Constructor(3)
	// obj.Put(1, 2)
	// obj.Put(3, 4)
	// obj.Put(5, 6)
	// log.Println(obj.Get(-1))
	// log.Println(obj.Get(1))
	// log.Println(obj.Get(3))
	// log.Println(obj.Get(5))

	obj := Constructor(10)
	obj.Put(10, 13)
	obj.Put(3, 17)
	obj.Put(6, 11)
	obj.Put(10, 5)
	obj.Put(9, 10)
	obj.Get(13)
	obj.Put(2, 19)
	obj.Get(2)
	obj.Get(3)
	obj.Put(5, 25)
	obj.Get(8)
	obj.Put(9, 22)
	obj.Put(5, 5)
	obj.Put(1, 30)
	obj.Get(11)
	obj.Put(9, 12)
	obj.Get(7)
	obj.Get(5)
	obj.Get(8)
	obj.Get(9)
	obj.Put(4, 30)
	obj.Put(9, 3)
	obj.Get(9)
	obj.Get(10)
	obj.Get(10)
	obj.Put(6, 14)
	obj.Put(3, 1)
	obj.Get(3)
	obj.Put(10, 11)
	obj.Get(8)
	obj.Put(2, 14)
	obj.Get(1)
	obj.Get(5)
	obj.Get(4)
	obj.Put(11, 4)
	obj.Put(12, 24)
	obj.Put(5, 18)
	obj.Get(13)
	obj.Put(7, 23)
	obj.Get(8)
	obj.Get(12)
	obj.Put(3, 27)
	obj.Put(2, 12)
	obj.Get(5)
	obj.Put(2, 9)
	obj.Put(13, 4)
	obj.Put(8, 18)
	obj.Put(1, 7)
	obj.Get(6)
	obj.Put(9, 29)
	obj.Put(8, 21)
	obj.Get(5)
	obj.Put(6, 30)
	obj.Put(1, 12)
	obj.Get(10)
	// obj.Put(9, 29)
	// obj.Put(8, 21)
	// obj.Get(5)
	// obj.Put(6, 30)
	// obj.Put(1, 12)
	// obj.Get(10)
	// obj.Put(4, 15)
	// obj.Put(7, 22)
	// obj.Put(11, 26)
	// obj.Put(8, 17)
	// obj.Put(9, 29)
}
