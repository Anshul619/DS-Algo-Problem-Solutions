package main

/*
- LeetCode - https://leetcode.com/problems/task-scheduler/
*/
import (
	"container/heap"
	"log"
)

type task struct {
	name      int
	frequency int
}

type tasksPriorityQueue []task

func leastInterval(tasks []byte, n int) int {

	pq := new(tasksPriorityQueue)
	dictionary := make([]int, 26)

	for _, v := range tasks {
		dictionary[v-'A']++
	}

	//log.Println(dictionary)
	for i, v := range dictionary {
		if v > 0 {
			temp := task{name: 'A' + i, frequency: v}
			heap.Push(pq, temp)
		}
	}

	//log.Println(pq)

	tasksBucket := make([]task, n+1)
	output := 0
	idle := 0

	for pq.Len() > 0 {

		idle = 0

		log.Println("pq before -", pq)

		for i := 0; i < len(tasksBucket); i++ {
			if pq.Len() > 0 {
				tasksBucket[i] = heap.Pop(pq).(task)
				tasksBucket[i].frequency--
				output++
			} else {
				log.Println("idle")
				idle++
			}
		}

		log.Println("tasksBucket -", tasksBucket)

		for _, t := range tasksBucket {
			if t.frequency > 0 {
				heap.Push(pq, t)
			}
		}

		if pq.Len() == 0 {
			break
		}

		output += idle
		//tasksBucket = tasksBucket[:0]
		//log.Println("pq after -", pq)
		log.Println("output -", output)
		log.Println("----")
	}

	return output
}

func (h tasksPriorityQueue) Len() int {
	return len(h)
}

func (h tasksPriorityQueue) Less(i int, j int) bool {
	return h[i].frequency > h[j].frequency
}

func (h tasksPriorityQueue) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *tasksPriorityQueue) Push(a interface{}) {
	*h = append(*h, a.(task))
}

func (h *tasksPriorityQueue) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func main() {

	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2

	//tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	//n := 0

	//tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	//n := 2
	log.Println(leastInterval(tasks, n))
}
