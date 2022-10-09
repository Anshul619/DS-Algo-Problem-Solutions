package main

/*
- LeetCode - https://leetcode.com/problems/queens-that-can-attack-the-king/submissions/
*/
import "log"

func calculateDistance(source []int, destination []int) int {

	distance := 0

	if destination[0] < source[0] {
		distance += source[0] - destination[0]
	} else {
		distance += destination[0] - source[0]
	}

	if destination[1] < source[1] {
		distance += source[1] - destination[1]
	} else {
		distance += destination[1] - source[1]
	}

	return distance
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {

	attackers := [8][]int{}
	out := [][]int{}

	for i := range attackers {
		attackers[i] = make([]int, 2)
		attackers[i][0] = -1
	}

	for _, v := range queens {

		if v[0] == king[0] && v[1] < king[1] {

			if attackers[3][0] == -1 || (calculateDistance(attackers[3], king) > calculateDistance(v, king)) {
				attackers[3] = v
			}
		}

		if v[0] == king[0] && v[1] > king[1] {

			if attackers[4][0] == -1 || (calculateDistance(attackers[4], king) > calculateDistance(v, king)) {
				attackers[4] = v
			}
		}

		// Attacker diagonally left & above the king
		if (king[0]-v[0] == king[1]-v[1]) && (king[0] > v[0]) {

			log.Println("Diagonally Left & Above the king")
			if attackers[0][0] == -1 || (calculateDistance(attackers[0], king) > calculateDistance(v, king)) {
				attackers[0] = v
			}
		}

		// Attacker diagonally left & below the king
		if (v[0]-king[0] == king[1]-v[1]) && (v[0] > king[0]) {

			log.Println("Diagonally Left & Below the king")
			if attackers[7][0] == -1 || (calculateDistance(attackers[7], king) > calculateDistance(v, king)) {
				attackers[7] = v
			}
		}

		if king[1] == v[1] && king[0] > v[0] {
			if attackers[1][0] == -1 || (calculateDistance(attackers[1], king) > calculateDistance(v, king)) {
				attackers[1] = v
			}
		}

		if king[1] == v[1] && king[0] < v[0] {
			if attackers[6][0] == -1 || (calculateDistance(attackers[6], king) > calculateDistance(v, king)) {
				attackers[6] = v
			}
		}

		// Attacker diagonally right & above the king
		if (king[0]-v[0] == v[1]-king[1]) && (v[0] < king[0]) {

			log.Println("Diagonally Right & Above the king")
			if attackers[2][0] == -1 || (calculateDistance(attackers[2], king) > calculateDistance(v, king)) {
				attackers[2] = v
			}
		}

		// Attacker diagonally right & below the king
		if (v[0]-king[0] == v[1]-king[1]) && (v[0] > king[0]) {

			log.Println("Diagonally Right & Below the king")
			if attackers[5][0] == -1 || (calculateDistance(attackers[5], king) > calculateDistance(v, king)) {
				attackers[5] = v
			}
		}
	}

	for _, v := range attackers {

		if v[0] != -1 {
			out = append(out, v)
		}
	}

	return out
}

func main() {
	queens := [][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}
	king := []int{0, 0}

	// queens := [][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}}

	// king := []int{3, 4}
	log.Println(queensAttacktheKing(queens, king))

}
