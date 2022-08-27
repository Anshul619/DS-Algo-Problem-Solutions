package main

//import "log"

type Queue []int

func (s *Queue) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Queue) Push(num int) {
	*s = append(*s, num)
}
func (s *Queue) Len() int {
	return len(*s)
}
func (s *Queue) Pop() (int, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		elem := (*s)[0]
		*s = (*s)[1:]
		return elem, true
	}
}

// func main() {
// 	queue := new(Queue)
// 	queue.Push(1)
// 	queue.Push(2)
// 	queue.Push(3)
// 	log.Println(queue.Pop())
// 	log.Println(queue.Pop())
// }
