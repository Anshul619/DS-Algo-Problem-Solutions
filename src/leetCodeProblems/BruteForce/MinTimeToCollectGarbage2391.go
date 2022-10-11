package main

/*
- LeetCode - https://leetcode.com/problems/minimum-amount-of-time-to-collect-garbage/
*/
import "log"

func garbageCollection(garbage []string, travel []int) int {

	MMax := -1
	GMax := -1
	PMax := -1

	out := 0

	for i, v := range garbage {

		Mfound := false
		Gfound := false
		Pfound := false

		for _, v1 := range v {
			rune := string(v1)

			if rune == "M" {
				Mfound = true
			} else if rune == "G" {
				Gfound = true
			} else if rune == "P" {
				Pfound = true
			}

			out++
		}

		if Mfound {
			MMax = i
		}

		if Gfound {
			GMax = i
		}

		if Pfound {
			PMax = i
		}

	}

	// log.Println(MMax, GMax, PMax)
	// //log.Println(out)
	for i, v := range travel {

		if i < MMax {
			out += v
		}

		if i < GMax {
			out += v
		}

		if i < PMax {
			out += v
		}

	}

	return out
}

func main() {
	// garbage := []string{"G", "P", "GP", "GG"}
	// travel := []int{2, 4, 3}

	garbage := []string{"MMM", "PGM", "GP"}
	travel := []int{3, 10}
	log.Println(garbageCollection(garbage, travel))

}
