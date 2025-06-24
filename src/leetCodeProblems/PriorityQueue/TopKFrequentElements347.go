package main

/*
- LeetCode - https://leetcode.com/problems/top-k-frequent-words/submissions/
- Time - O(N) + O(N log K) = O(N log K)
- Space - O(N + K)
*/
import (
	"container/heap"
)

type Element struct {
	Num  int
	Freq int
}

type minHeap []Element

// Min Heap
func (h minHeap) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h *minHeap) Push(i interface{}) {
	*h = append(*h, i.(Element))
}

func (h *minHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)
	h := new(minHeap)
	heap.Init(h)

	for _, v := range nums {
		m[v]++
	}

	for n, f := range m {
		if h.Len() >= k {
			e := heap.Pop(h).(Element)
			if e.Freq < f {
				heap.Push(h, Element{n, f})
			} else {
				heap.Push(h, e)
			}
		} else {
			heap.Push(h, Element{n, f})
		}
	}

	out := []int{}

	for i := 0; i < k; i++ {
		out = append(out, heap.Pop(h).(Element).Num)
	}
	return out
}
