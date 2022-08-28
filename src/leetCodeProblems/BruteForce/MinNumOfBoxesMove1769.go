package main

import "log"

func minOperations(boxes string) []int {

	out := make([]int, len(boxes))

	balls := []int{}

	for i, v := range boxes {
		// log.Println()
		// log.Println(boxes[i])
		if string(v) == "1" {
			balls = append(balls, i)
		}
	}

	// log.Println(balls)

	for i := range out {

		move := 0

		for _, j := range balls {

			if i > j {
				move += i - j
			} else {
				move += j - i
			}
		}

		out[i] = move
	}
	return out
}

func main() {
	//boxes := "110"
	boxes := "001011"
	log.Println(minOperations(boxes))
}
