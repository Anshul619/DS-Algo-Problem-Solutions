package main

/*
PayPay OA (Code Signal) - Question 1
- Time - O(n)
- Space - O(1)

Given an array of integers numbers, find the first index i for which three consecutive elements numbers[i], numbers[i + 1], and numbers[i + 2] are all greater than a given threshold. If there is no such index, return -1.
Note: You are not expected to provide the most optimal solution, but a solution with time complexity not worse than O(numbers.length2) will fit within the execution time limit.
Example
* For numbers = [0, 1, 4, 3, 2, 5], and threshold = 1 the output should be solution(numbers, threshold) = 2.Explanation:
  - When i = 0, the triple (numbers[0], numbers[1], numbers[2]) = (0, 1, 4), contains 0 and 1 which are not greater than threshold = 1;
  - When i = 1, the triple (numbers[1], numbers[2], numbers[3]) = (1, 4, 3) contains 1 which is not greater than threshold = 1;
  - When i = 2, the triple (numbers[2], numbers[3], numbers[4]) = (4, 3, 2) contains all numbers which are greater than threshold = 1;
  - Thus, the answer is i = 2.

* For numbers = [-9, 95, 94, 4, 51], and threshold = 42 the output should be solution(numbers, threshold) = -1.Explanation:
  - When i = 0, the triple (numbers[0], numbers[1], numbers[2]) = (-9, 95, 94) contains -9 < 42 which does not satisfy the condition;
  - When i = 1, the triple (numbers[1], numbers[2], numbers[3]) = (95, 94, 4), contains 4 < 42 which does not satisfy the condition;
  - When i = 2, the triple (numbers[2], numbers[3], numbers[4]) = (94, 4, 51), contains 4 < 42 which does not satisfy the condition;
  - None of the triples satisfy the condition, so the answer is -1.
*/
func solution1(numbers []int, threshold int) int {
	for i := 0; i < len(numbers)-2; i++ {
		if numbers[i] > threshold &&
			numbers[i+1] > threshold &&
			numbers[i+2] > threshold {
			return i
		}
	}
	return -1
}
