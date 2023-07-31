package main

/*
- Leetcode - https://leetcode.com/problems/reorganize-string/description/
- Time - O(nlogn)
- Space - O(n)
*/
import (
	"container/heap"
)

type Node struct {
	char  rune
	count int
}

type CharPriorityQueue []Node

func (pq CharPriorityQueue) Len() int {
	return len(pq)
}

// Max-heap
func (pq CharPriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq CharPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *CharPriorityQueue) Push(num interface{}) {
	*pq = append(*pq, num.(Node))
}

func (pq *CharPriorityQueue) Pop() interface{} {
	len := len(*pq)
	num := (*pq)[len-1]
	*pq = (*pq)[:len-1]
	return num
}

func reorganizeString(s string) string {
	countsMap := [26]int{}
	out := []rune{}

	for _, v := range s {
		countsMap[v-'a']++
	}

	pq := new(CharPriorityQueue)

	for i, v := range countsMap {
		if v > 0 {
			heap.Push(pq, Node{rune(i + 'a'), v})
		}
	}

	var prev Node

	for pq.Len() > 0 {
		node := heap.Pop(pq).(Node)
		out = append(out, node.char)

		if prev.count > 0 {
			heap.Push(pq, prev)
		}

		node.count--
		prev = node
	}

	if prev.count > 0 {
		return ""
	}

	return string(out)
}

// func main() {
// 	log.Println(reorganizeString("a"))
// 	log.Println(reorganizeString("aab"))
// 	log.Println(reorganizeString("aaab"))
// 	log.Println(reorganizeString("aaabc"))
// 	log.Println(reorganizeString("aaabb"))
// }
