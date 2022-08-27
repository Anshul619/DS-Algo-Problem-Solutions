package main

/*
- LeetCode - https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to
*/
import "fmt"

func addNewGroup(elem int, size int, output *[][]int, m map[int]int) {
	temp := []int{}
	temp = append(temp, elem)

	*output = append(*output, temp)
	m[size] = len(*output) - 1
}

func groupThePeople(groupSizes []int) [][]int {

	output := [][]int{}
	m := make(map[int]int)

	for elem, size := range groupSizes {
		if outIndex, ok := m[size]; ok {

			if len(output[outIndex]) < size {
				output[outIndex] = append(output[outIndex], elem)
			} else {
				addNewGroup(elem, size, &output, m)
			}

		} else {
			addNewGroup(elem, size, &output, m)
		}
	}

	return output
}

func main() {

	input := []int{3, 3, 3, 3, 3, 1, 3}
	//input := []int{2, 1, 3, 3, 3, 2}
	fmt.Println(groupThePeople(input))
}
