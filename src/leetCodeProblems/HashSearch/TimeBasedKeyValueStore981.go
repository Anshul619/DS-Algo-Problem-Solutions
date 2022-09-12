package main

/*
- LeetCode - https://leetcode.com/problems/time-based-key-value-store/solution/
*/

import (
	"log"
	"strconv"
)

type PriorityQueue []int

func (q *PriorityQueue) push(e int) {
	*q = append(*q, e)
}

func (q *PriorityQueue) pop() int {
	e := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return e
}

func (q PriorityQueue) Less(i int, j int) bool {
	return q[i] > q[j]
}

func (q PriorityQueue) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q PriorityQueue) Len() int {
	return len(q)
}

type stack []int

func (s *stack) pop() int {
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *stack) push(e int) {
	*s = append(*s, e)
}

func (s stack) isEmpty() bool {
	return len(s) == 0
}

type TimeMap struct {
	m map[string]string
	q *PriorityQueue
}

func Constructor() TimeMap {
	tm := new(TimeMap)
	tm.m = make(map[string]string)
	tm.q = new(PriorityQueue)
	return *tm
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	mKey := key + "_" + strconv.Itoa(timestamp)
	this.m[mKey] = value
	this.q.push(timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {

	mKey := key + "_" + strconv.Itoa(timestamp)
	tStack := new(stack)
	out := "test"

	if v, ok := this.m[mKey]; ok {
		return v
	}

	for this.q.Len() > 0 {

		e := this.q.pop()
		mKey1 := key + "_" + strconv.Itoa(e)

		tStack.push(e)

		if e < timestamp {

			if v, ok := this.m[mKey1]; ok {
				out = v
				break
			}

		}
	}

	for !tStack.isEmpty() {
		this.q.push(tStack.pop())
	}

	return out
}

func main() {
	// tm := Constructor()
	// tm.Set("foo", "bar", 1)
	// log.Println(tm.Get("foo", 1))
	// log.Println(tm.Get("foo", 3))
	// tm.Set("foo", "bar2", 4)
	// log.Println(tm.Get("foo", 4))
	// log.Println(tm.Get("foo", 5))

	tm := Constructor()
	tm.Set("ctondw", "ztpearaw", 1)
	tm.Set("vrobykydll", "hwliiq", 2)
	tm.Set("gszaw", "ztpearaw", 3)
	tm.Set("ctondw", "gszaw", 4)
	log.Println(tm.Get("gszaw", 5))

}
