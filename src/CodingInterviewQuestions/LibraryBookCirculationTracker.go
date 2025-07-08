package main

import (
	"sort"
	"strconv"
)

/*

PayPay OA (Code Signal) - Question 3
- Time - O(m · k · log k)
- Space - O(1)

Your task is to develop a library book circulation tracker.
You are given a sequence of operations that represent activities in a library. Each operation is one of three types: acquisition, checkout, or reclassify. Operations are provided in the following format:
* ["acquisition", "<book category>", "<quantity>", "<price>"] – the library acquires <quantity> books of <book category>, each valued at <price> for insurance purposes.
* ["checkout", "<book category>", "<quantity>"] – patrons borrow <quantity> books of <book category>. If books of the specified category have different insurance values, the least valuable ones should be checked out first. It is guaranteed that the library will always have enough books to fulfill all checkout requests.
* ["reclassify", "<book category>", "<quantity>", "<original price>", "<new price>"] – the library reclassifies <quantity> books of <book category> to a more valuable edition. It is guaranteed that there are <quantity> books of the specified category with the <original price> value.
Your function should calculate the total insurance value of all books checked out after processing the entire sequence of operations. Return an array representing the insurance value of books for each checkout operation.
Note: You are not expected to provide the most optimal solution, but a solution with time complexity not worse than O(operations.length2) will fit within the execution time limit.
Example
For
operations = [
  ["acquisition", "fiction", "2", "100"],
  ["acquisition", "reference", "3", "60"],
  ["checkout", "fiction", "1"],
  ["checkout", "reference", "1"],
  ["reclassify", "reference", "1", "60", "100"],
  ["checkout", "reference", "1"],
  ["checkout", "reference", "1"]
]
the output should be
solution(operations) = [100, 60, 60, 100]
Let's analyze the operations:
* 		["acquisition", "fiction", "2", "100"] - the library acquires 2 fiction books, each valued at 100.
* 		["acquisition", "reference", "3", "60"] - the library acquires 3 reference books, each valued at 60.
* 		["checkout", "fiction", "1"] - a patron checks out 1 fiction book valued at 100, so the insurance value is 1 × 100 = 100.
* 		["checkout", "reference", "1"] - a patron checks out 1 reference book valued at 60, so the insurance value is 1 × 60 = 60.
* 		["reclassify", "reference", "1", "60", "100"] - one reference book is reclassified and its value becomes 100.
* 		["checkout", "reference", "1"] - a patron checks out 1 reference book. There are 2 reference books in the library with values of 60 and 100, and the one with the value of 60 should be checked out first, so the insurance value is 1 × 60 = 60.
* 		["checkout", "reference", "1"] - a patron checks out 1 reference book. There is 1 reference book in the library with the value of 100, so the insurance value is 1 × 100 = 100.
*/

func solution3(operations [][]string) []int {
	//books[category][price] = quantity
	m := make(map[string]map[int]int)
	result := []int{}

	for _, op := range operations {
		switch op[0] {
		case "acquisition":
			category := op[1]
			quantity, _ := strconv.Atoi(op[2])
			price, _ := strconv.Atoi(op[3])
			if m[category] == nil {
				m[category] = make(map[int]int)
			}
			m[category][price] += quantity

		case "checkout":
			category := op[1]
			quantity, _ := strconv.Atoi(op[2])
			priceMap := m[category]

			prices := []int{}
			for price := range priceMap {
				prices = append(prices, price)
			}
			sort.Ints(prices)

			insuranceValue := 0
			remaining := quantity

			for _, price := range prices {
				count := priceMap[price]
				if count == 0 {
					continue
				}

				// find min value
				toRemove := min(remaining, count)
				insuranceValue += toRemove * price
				priceMap[price] -= toRemove
				remaining -= toRemove
				if remaining == 0 {
					break
				}
			}

			result = append(result, insuranceValue)

		case "reclassify":
			category := op[1]
			quantity, _ := strconv.Atoi(op[2])
			originalPrice, _ := strconv.Atoi(op[3])
			newPrice, _ := strconv.Atoi(op[4])
			priceMap := m[category]

			priceMap[originalPrice] -= quantity
			priceMap[newPrice] += quantity

		}
	}
	return result
}
