package main

/*
- LeetCode - https://leetcode.com/problems/cheapest-flights-within-k-stops
*/
import (
	"math"
)

type Connection struct {
	city int
	cost int
}

type queue []Connection

func (q *queue) push(c Connection) {
	q1 := []Connection{c}
	*q = append(*q, q1...)
}

func (q *queue) pop() (bool, Connection) {

	if q.isEmpty() {
		return false, Connection{}
	} else {
		elem := (*q)[0]
		*q = (*q)[1:]
		return true, elem
	}
}

func (q queue) isEmpty() bool {
	return q.size() <= 0
}

func (q queue) size() int {
	return len(q)
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {

	minCost := math.MaxInt64

	fQueue := new(queue)
	fQueue.push(Connection{src, 0})

	maxStops := k + 1

	costOfFlights := make([][]int, n)
	for i := range costOfFlights {
		costOfFlights[i] = make([]int, n)
	}

	flightsMap := make(map[int][]int)
	for _, v := range flights {
		if _, ok := flightsMap[v[0]]; ok {
			flightsMap[v[0]] = append(flightsMap[v[0]], v[1])
		} else {
			flightsMap[v[0]] = []int{v[1]}
		}

		costOfFlights[v[0]][v[1]] = v[2]
	}

	cityPathCost := make([]int, n) // path cost from source
	for i := range cityPathCost {
		cityPathCost[i] = math.MaxInt64
	}

	cityPathCost[src] = 0

	for !fQueue.isEmpty() && maxStops >= 0 {

		size := fQueue.size()
		maxStops--

		for size > 0 {
			_, c := fQueue.pop()
			size--

			if c.city == dst {
				if minCost > c.cost {
					minCost = c.cost
				}
				continue
			}

			for _, v := range flightsMap[c.city] {

				newCost := c.cost + costOfFlights[c.city][v]

				if cityPathCost[v] > newCost {
					fQueue.push(Connection{v, newCost})
					cityPathCost[v] = newCost
				}
			}
		}
	}

	if minCost == math.MaxInt64 {
		minCost = -1
	}

	return minCost
}

// func main() {
// 	//input := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
// 	//log.Println(findCheapestPrice(4, input, 0, 3, 1))

// 	// input := [][]int{{0, 1, 100}, {1, 2, 100}}
// 	// log.Println(findCheapestPrice(3, input, 0, 2, 1))

// 	//input := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
// 	//log.Println(findCheapestPrice(3, input, 0, 2, 0))

// 	//input := [][]int{{3, 4, 4}, {2, 5, 6}, {4, 7, 10}, {9, 6, 5}, {7, 4, 4}, {6, 2, 10}, {6, 8, 6}, {7, 9, 4}, {1, 5, 4}, {1, 0, 4}, {9, 7, 3}, {7, 0, 5}, {6, 5, 8}, {1, 7, 6}, {4, 0, 9}, {5, 9, 1}, {8, 7, 3}, {1, 2, 6}, {4, 1, 5}, {5, 2, 4}, {1, 9, 1}, {7, 8, 10}, {0, 4, 2}, {7, 2, 8}}
// 	//log.Println(findCheapestPrice(10, input, 6, 0, 7))

// 	input := [][]int{{1, 2, 10}, {2, 0, 7}, {1, 3, 8}, {4, 0, 10}, {3, 4, 2}, {4, 2, 10}, {0, 3, 3}, {3, 1, 6}, {2, 4, 5}}
// 	log.Println(findCheapestPrice(5, input, 0, 4, 1))
// }
