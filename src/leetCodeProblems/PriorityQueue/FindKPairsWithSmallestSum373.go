package main

/*
- Leetcode - https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
- Time - O(min(k*logk),m * n * log(m * n))
- Space - O(min(k, m * n))
*/
import (
	"container/heap"
)

type Pair struct {
	index1, index2 int
}

type PairSum struct {
	sum, index1, index2 int
}

type pairQueue []PairSum

// Min-Heap
func (pq pairQueue) Less(i, j int) bool {
	return pq[i].sum < pq[j].sum
}

func (pq pairQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pairQueue) Push(p interface{}) {
	(*pq) = append((*pq), p.(PairSum))
}

func (pq *pairQueue) Pop() interface{} {
	t := (*pq)[pq.Len()-1]
	*pq = (*pq)[:pq.Len()-1]
	return t
}

func (pq pairQueue) isEmpty() bool {
	return len(pq) == 0
}

func (pq pairQueue) Len() int {
	return len(pq)
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	pq := new(pairQueue)
	visited := make(map[Pair]bool)

	heap.Push(pq, PairSum{nums1[0] + nums2[0], 0, 0})
	visited[Pair{0, 0}] = true

	out := [][]int{}

	for k > 0 && !pq.isEmpty() {
		p := heap.Pop(pq).(PairSum)
		out = append(out, []int{nums1[p.index1], nums2[p.index2]})

		if p.index1+1 < len(nums1) && !visited[Pair{p.index1 + 1, p.index2}] {
			heap.Push(pq, PairSum{nums1[p.index1+1] + nums2[p.index2], p.index1 + 1, p.index2})
			visited[Pair{p.index1 + 1, p.index2}] = true
		}

		if p.index2+1 < len(nums2) && !visited[Pair{p.index1, p.index2 + 1}] {
			heap.Push(pq, PairSum{nums1[p.index1] + nums2[p.index2+1], p.index1, p.index2 + 1})
			visited[Pair{p.index1, p.index2 + 1}] = true
		}
		k--
	}
	return out
}
