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
		return
	} else if keyToBeRemoved == 0 {
		*s = (*s)[1:]
	} else if keyToBeRemoved == len(*s)-1 {
		*s = (*s)[:len(*s)-2]
	} else {
		*s = append((*s)[:keyToBeRemoved-1], (*s)[keyToBeRemoved+1:]...)
	}

	return
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
	this.lruQueue.RemoveKey(key)

	if this.size == this.lruQueue.Len() {
		//fmt.Println("REMOVe")
		this.lruQueue.Remove()
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

// func main() {
// 	obj := Constructor(3)
// 	obj.Put(1, 2)
// 	obj.Put(3, 4)
// 	obj.Put(5, 6)
// 	log.Println(obj.Get(-1))
// 	log.Println(obj.Get(1))
// 	log.Println(obj.Get(3))
// 	log.Println(obj.Get(5))
// }
