package main

import "testing"

func TestLRUCache(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		lru := Constructor(2)
		lru.Put(1, 1)
		lru.Put(2, 2)
		if lru.Get(1) != 1 {
			t.Fatalf("get 1")
		}
		lru.Put(3, 3)
		if lru.Get(2) != -1 {
			t.Fatalf("get 2")
		}
		lru.Put(4, 4)
		if lru.Get(1) != -1 {
			t.Fatalf("get 1 after put 4")
		}
		if lru.Get(3) != 3 {
			t.Fatalf("get 3")
		}
		if lru.Get(4) != 4 {
			t.Fatalf("get 4")
		}
	})

	t.Run("test2", func(t *testing.T) {
		lru := Constructor(1)
		lru.Put(2, 1)
		if lru.Get(2) != 1 {
			t.Fail()
		}
	})
}
