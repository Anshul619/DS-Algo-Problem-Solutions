package main

/*
- LeetCode - https://leetcode.com/problems/top-k-frequent-words/submissions/
- Time - O(N) + O(N log K) + O(K log K) = O(N log K)
- Space - O(N + K)
*/
import (
	"container/heap"
)

type Element struct {
	num  int
	freq int
}

type minHeap []Element

// Sort interface implementation

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Pop() interface{} {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func (h *minHeap) Push(t interface{}) {
	*h = append(*h, t.(Element))
}

func topKFrequent1(nums []int, k int) []int {
	pq := new(minHeap)
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for num, freq := range m {
		if pq.Len() < k {
			heap.Push(pq, Element{num, freq})
		} else if (*pq)[0].freq < freq {
			heap.Pop(pq)
			heap.Push(pq, Element{num, freq})
		}
	}

	out := make([]int, k)

	for i := 0; i < k; i++ {
		out[i] = heap.Pop(pq).(Element).num
	}
	return out
}
