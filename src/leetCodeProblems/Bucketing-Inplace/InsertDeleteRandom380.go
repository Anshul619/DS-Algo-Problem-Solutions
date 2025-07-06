package main

/*
- LeetCode - https://leetcode.com/problems/insert-delete-getrandom-o1/
- Space - O(n)
- TIme - O(1)
*/
import "math/rand"

type RandomizedSet struct {
	m    map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	r := RandomizedSet{}
	r.m = make(map[int]int)
	r.nums = []int{}
	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	lastIndex := len(this.nums) - 1
	lastVal := this.nums[lastIndex]

	// swap elements
	this.nums[this.m[val]] = lastVal
	this.m[lastVal] = this.m[val]

	this.nums = this.nums[:lastIndex]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	randIndex := rand.Intn(len(this.nums))
	return this.nums[randIndex]
}
