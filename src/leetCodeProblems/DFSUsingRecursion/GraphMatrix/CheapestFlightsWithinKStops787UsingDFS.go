package main

/*
LeetCode - https://leetcode.com/problems/cheapest-flights-within-k-stops
*/

type ConnectionDFS struct {
	endCity int
	cost    int
}

func buildMapDFS(flights [][]int) map[int][]ConnectionDFS {
	flightsMap := make(map[int][]ConnectionDFS)

	for _, v := range flights {

		routeNode := ConnectionDFS{v[1], v[2]}

		if _, ok := flightsMap[v[0]]; ok {
			flightsMap[v[0]] = append(flightsMap[v[0]], routeNode)
		} else {
			flightsMap[v[0]] = []ConnectionDFS{routeNode}
		}

	}

	return flightsMap
}

func findCheapestPriceDFSUtil(cityRoute ConnectionDFS, flightsMap map[int][]ConnectionDFS, src int, dst int, maxStops int, routeCost int, minCost int, path []int) int {

	if cityRoute.endCity == src && len(path) > 0 {
		return minCost
	}

	if minCost != -1 && routeCost > minCost {
		return minCost
	}

	routeCost += cityRoute.cost
	path = append(path, cityRoute.endCity)

	if cityRoute.endCity == dst {
		if minCost == -1 || routeCost < minCost {
			return routeCost
		} else {
			return minCost
		}
	}

	if cityRoute.endCity != src {
		maxStops--
	}

	if maxStops < 0 {
		return minCost
	}

	for _, v := range flightsMap[cityRoute.endCity] {
		minCost = findCheapestPriceDFSUtil(v, flightsMap, src, dst, maxStops, routeCost, minCost, path)
	}

	return minCost

}

func findCheapestPriceDFS(n int, flights [][]int, src int, dst int, k int) int {

	flightsMap := buildMapDFS(flights)

	routeCities := []int{}
	node := ConnectionDFS{src, 0}

	return findCheapestPriceDFSUtil(node, flightsMap, src, dst, k, 0, -1, routeCities)
}

// func main() {
// 	//input := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
// 	//log.Println(findCheapestPrice(4, input, 0, 3, 1))

// 	// input := [][]int{{0, 1, 100}, {1, 2, 100}}
// 	// log.Println(findCheapestPrice(3, input, 0, 2, 1))

// 	input := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
// 	log.Println(findCheapestPriceDFS(3, input, 0, 2, 0))
// }
