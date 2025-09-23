package main

/*
- Leetcode - https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/description/
- Time - O(n)
- Space - O(n)
*/

import (
	"container/heap"
)

type Character struct {
	char      rune
	frequency int
}
type charPQ []Character

func (h charPQ) Len() int {
	return len(h)
}

// Max-heap
func (h charPQ) Less(i int, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h charPQ) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *charPQ) Push(a any) {
	*h = append(*h, a.(Character))
}

func (h *charPQ) Pop() any {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func minDeletions(s string) int {
	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	pq := new(charPQ)

	for k, v := range m {
		heap.Push(pq, Character{k, v})
	}

	out := 0

	for len(*pq) > 1 {
		e1, e2 := heap.Pop(pq).(Character), heap.Pop(pq).(Character)

		if e1.frequency != e2.frequency {
			heap.Push(pq, e2)
		} else {
			if e1.frequency > 1 {
				heap.Push(pq, Character{e1.char, e1.frequency - 1})
			}

			heap.Push(pq, e2)
			out++
		}
	}
	return out
}
