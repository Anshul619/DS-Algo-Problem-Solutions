package main

/*
- Leetcode - https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/description/
- Time - O(n)
- Space - O(n)
*/
import (
	"container/heap"
)

type CharCount struct {
	char  rune
	count int
}

type CharCountPQ []CharCount

func (pq CharCountPQ) Len() int {
	return len(pq)
}

// Max-heap
func (pq CharCountPQ) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq CharCountPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *CharCountPQ) Push(node interface{}) {
	*pq = append(*pq, node.(CharCount))
}

func (pq *CharCountPQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return temp
}

func (pq CharCountPQ) isEmpty() bool {
	return len(pq) == 0
}

func minDeletions(s string) int {
	m := make(map[rune]int)
	pq := new(CharCountPQ)

	out := 0

	for _, v := range s {
		m[v]++
	}

	for k, v := range m {
		heap.Push(pq, CharCount{k, v})
	}

	for !pq.isEmpty() {
		elem := heap.Pop(pq).(CharCount)

		if !pq.isEmpty() {

			peek := heap.Pop(pq).(CharCount)
			if peek.count == elem.count {
				elem.count--

				if elem.count > 0 {
					heap.Push(pq, elem)
				}

				heap.Push(pq, peek)
				out++
			} else {
				heap.Push(pq, peek)
			}
		}
	}
	return out
}

// func main() {
// 	// log.Println(minDeletions("ceabaacb"))
// 	// log.Println(minDeletions("aab"))
// 	// log.Println(minDeletions("aaabbbcc"))
// 	log.Println(minDeletions("bbcebab"))
// }
