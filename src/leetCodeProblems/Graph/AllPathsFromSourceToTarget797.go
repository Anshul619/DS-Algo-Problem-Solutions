package main

func util(graph [][]int, temp *[]int, out *[][]int, num int, target int) {

	*temp = append(*temp, num)

	if num == target {
		*out = append(*out, *temp)
		return
	}

	for _, v := range graph[num] {

		temp1 := []int{}
		temp1 = append(temp1, *temp...)
		util(graph, &temp1, out, v, target)
	}

}
func allPathsSourceTarget(graph [][]int) [][]int {

	out := [][]int{}
	target := len(graph) - 1

	for _, v := range graph[0] {
		temp := []int{0}
		util(graph, &temp, &out, v, target)
	}

	return out
}

// func main() {

// 	graph := [][]int{{1, 2}, {3}, {3}, {}}
// 	//graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}

// 	log.Println(allPathsSourceTarget(graph))
// }
