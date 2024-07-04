package main

import (
	"testing"
)

func TestPeekingIterator(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		iter := new(Iterator)
		iter.nums = []int{1, 2, 3}

		pIter := Constructor(iter)
		pIter.iter = iter

		if pIter.next() != 1 {
			t.Errorf("failed")
		}

		if pIter.peek() != 2 {
			t.Errorf("failed")
		}

		if pIter.next() != 2 {
			t.Errorf("failed")
		}

		if pIter.next() != 3 {
			t.Errorf("failed")
		}

		if pIter.hasNext() {
			t.Errorf("failed")
		}
	})
}
