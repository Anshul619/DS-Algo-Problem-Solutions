package main

import "log"

/*
- LeetCode - https://leetcode.com/problems/design-parking-system/
- Time - O(1)
- Space - O(1)
*/

type OrderedStream struct {
	stream []string
	ptr    int
}

func Constructor3(n int) OrderedStream {
	return OrderedStream{make([]string, n), 0}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.stream[idKey-1] = value
	oldPtr := this.ptr

	for this.ptr < len(this.stream) && this.stream[this.ptr] != "" {
		this.ptr++
	}

	return this.stream[oldPtr:this.ptr]
}

func main() {
	obj := Constructor3(5)

	log.Println(obj.Insert(3, "ccccc"))
	log.Println(obj.Insert(1, "aaaaa"))
	log.Println(obj.Insert(2, "bbbbb"))
	log.Println(obj.Insert(5, "eeeee"))
	log.Println(obj.Insert(4, "ddddd"))

	// log.Println(obj.Get(-1))
	// log.Println(obj.Get(1))
	// log.Println(obj.Get(3))
	// log.Println(obj.Get(5))
}
