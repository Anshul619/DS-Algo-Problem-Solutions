package main

/*
- LeetCode - https://leetcode.com/problems/top-k-frequent-words/submissions/
*/
import (
	"container/heap"
	//"log"
)

type wordFreq struct {
	word      string
	frequency int
}

type pq []wordFreq

func (h pq) Len() int {
	return len(h)
}

func (h *pq) Push(a interface{}) {
	*h = append(*h, a.(wordFreq))
}

func (h *pq) Pop() interface{} {
	len := len(*h)
	elem := (*h)[len-1]
	*h = (*h)[:len-1]
	return elem
}

func (h pq) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h pq) Less(i int, j int) bool {

	if h[i].frequency == h[j].frequency {
		return h[i].word < h[j].word
	}
	return h[i].frequency > h[j].frequency
}

func topKFrequent(words []string, k int) []string {

	out := make([]string, k)
	dic := make(map[string]int)
	pqObj := new(pq)

	for _, v := range words {
		if _, ok := dic[v]; ok {
			dic[v]++
		} else {
			dic[v] = 1
		}
	}

	for i, v := range dic {
		heap.Push(pqObj, wordFreq{word: i, frequency: v})
	}

	for i := 0; i < k; i++ {
		out[i] = heap.Pop(pqObj).(wordFreq).word
	}

	return out
}

// func main() {

// 	//input := []string{"i", "love", "leetcode", "i", "love", "coding"}
// 	//k := 2

// 	input := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
// 	k := 4

// 	//input := []string{"i", "love", "leetcode", "i", "love", "coding"}
// 	//k := 3
// 	log.Println(topKFrequent(input, k))
// }
