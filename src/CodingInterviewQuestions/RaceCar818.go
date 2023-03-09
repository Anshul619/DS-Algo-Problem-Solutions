package main

/*
- LeetCode - https://leetcode.com/problems/race-car/solution/
*/
import (
	"log"
	"math"
)

func util(dist int, ins *string) {

	if dist == 0 {
		return
	}

	if dist == 1 {
		*ins += "A"
		return
	}

	maxDistance := 0
	i := 0

	for maxDistance <= dist {

		maxDistance = int(math.Pow(2, float64(i)))

		if maxDistance <= dist {
			*ins += "A"
		}

		i++
	}

	*ins += "R"

	util(maxDistance-dist, ins)
}

func racecar(target int) int {

	inst := ""

	util(target, &inst)

	log.Println(inst)

	return len(inst)
}

// func main() {
// 	target := 3
// 	//target := 25

// 	racecar(target)
// }
