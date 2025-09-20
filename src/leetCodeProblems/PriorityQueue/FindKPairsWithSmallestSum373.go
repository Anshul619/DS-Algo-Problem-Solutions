package main

import "container/heap"

/*
- Leetcode - https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
- Time - O(min(n,k)+klog(min(n,k))) = O(klog(min(n,k)))
- Space - O(min(n, k))
*/
type Pair struct {
	i   int // index in Num1
	j   int // index in Num2
	sum int
}

type minHeap3 []Pair

func (h minHeap3) Len() int {
	return len(h)
}

func (h minHeap3) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h minHeap3) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap3) Push(x any) {
	*h = append(*h, x.(Pair))
}

func (h *minHeap3) Pop() any {
	t := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return t
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := new(minHeap3)
	heap.Init(h)

	// seed heap with (nums1[i], nums2[0]) for first min(len(nums1), k)
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Pair{i, 0, nums1[i] + nums2[0]})
	}

	out := [][]int{}

	for k > 0 && h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		out = append(out, []int{nums1[p.i], nums2[p.j]})
		k--

		// push next pair (i, j+1) if exists
		if p.j+1 < len(nums2) {
			heap.Push(h, Pair{p.i, p.j + 1, nums1[p.i] + nums2[p.j+1]})
		}
	}

	return out
}
