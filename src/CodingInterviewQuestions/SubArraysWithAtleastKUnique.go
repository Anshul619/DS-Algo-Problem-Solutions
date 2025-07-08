package main

/**
PayPay OA (Code Signal) - Question 4

Similar leetcode problem - https://leetcode.com/problems/subarrays-with-k-different-integers/description/?utm_source=chatgpt.com

Imagine you are an inventory manager at a warehouse filled with various items represented by an array inventory. The warehouse employs a logistics system that can identify the uniqueness of items within specified sections of the warehouse. Given these sections represented by contiguous subarrays, your task is to determine how many of these sections contain at least k unique types of items.
Example
* For inventory = [1, 2, 1, 1] and k = 2, the output should be solution(inventory, k) = 2.There are 2 sections that satisfy the condition for having at least k = 2 unique types of items:
    * inventory[0..1] = [1, 2]
    * inventory[1..2] = [2, 1]
* Note that the section inventory[0..2] = [1, 2, 1] is not counted because it contains the number 1 twice, and thus doesn't have 2 unique items.
* For inventory = [1, 2, 3, 4, 1] and k = 3, the output should be solution(inventory, k) = 6.There are 6 sections that contain at least k = 3 unique types of items:
    * inventory[0..2] = [1, 2, 3]
    * inventory[0..3] = [1, 2, 3, 4]
    * inventory[0..4] = [1, 2, 3, 4, 1] - even though there are two 1s, it still contains 3 unique items.
    * inventory[1..3] = [2, 3, 4]
    * inventory[1..4] = [2, 3, 4, 1]
    * inventory[2..4] = [3, 4, 1]
* For inventory = [5, 5, 5, 5] and k = 2, the output should be solution(inventory, k) = 0.No sections contain at least k = 2 unique items.
* For inventory = [5, 5, 5, 5] and k = 1, the output should be solution(inventory, k) = 4.There are 4 sections satisfying the condition with at least k = 1 unique items:
    * inventory[0..0] = [5]
    * inventory[1..1] = [5]
    * inventory[2..2] = [5]
    * inventory[3..3] = [5]
*/

func solution4(inventory []int, k int) int {
	n := len(inventory)
	count := 0
	freq := make(map[int]int)
	unique := 0
	left := 0

	for right := 0; right < n; right++ {
		item := inventory[right]
		if freq[item] == 0 {
			unique++
		}
		freq[item]++

		for unique >= k {
			// Count all subarrays starting at left and ending at or after right
			count += n - right
			// Try shrinking the window from the left
			freq[inventory[left]]--
			if freq[inventory[left]] == 0 {
				unique--
			}
			left++
		}
	}

	return count
}
