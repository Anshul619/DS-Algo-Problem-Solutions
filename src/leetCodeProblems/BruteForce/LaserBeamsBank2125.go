package main

import (
	"log"
	"strings"
)

func numberOfBeams(bank []string) int {

	out := 0
	previousCount := 0

	for _, row := range bank {

		currentCount := 0

		if !strings.Contains(row, "1") {
			continue
		}

		for _, cell := range row {

			if string(cell) == "1" {
				currentCount++
			}
		}

		// if previousCount != -1 {

		// }
		out += currentCount * previousCount
		previousCount = currentCount

	}
	return out
}

func main() {

	bank := []string{"011001", "000000", "010100", "001000"}
	//bank := []string{"000", "111", "000"}
	log.Println(numberOfBeams(bank))
}
