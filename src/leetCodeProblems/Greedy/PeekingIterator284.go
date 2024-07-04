package main

/*
- Leetcode - https://leetcode.com/problems/peeking-iterator/description/
- Time - O(1)
- Space - O(1)
*/
type Iterator struct {
	nums  []int
	index int
}

func (this *Iterator) hasNext() bool {
	return this.index < len(this.nums)
}

func (this *Iterator) next() int {
	n := this.nums[this.index]
	this.index++
	return n
}

func advanceIterator(pIter *PeekingIterator) {
	if pIter.iter.hasNext() {
		pIter.nextVal = pIter.iter.next()
	} else {
		pIter.isEmpty = true
	}
}

type PeekingIterator struct {
	iter    *Iterator
	nextVal int
	isEmpty bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	pIter := new(PeekingIterator)
	pIter.iter = iter
	advanceIterator(pIter)
	return pIter
}

func (this *PeekingIterator) hasNext() bool {
	return !this.isEmpty
}

func (this *PeekingIterator) next() int {
	nxt := this.nextVal
	advanceIterator(this)
	return nxt
}

func (this *PeekingIterator) peek() int {
	return this.nextVal
}
