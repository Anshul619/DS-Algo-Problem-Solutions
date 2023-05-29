package main

/*
- LeetCode - https://leetcode.com/problems/convert-the-temperature/solutions/3558115/o-1-time-simple-solution/
*/

func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32.00}
}

// func main() {
// 	log.Println(convertTemperature(36.50))
// }
