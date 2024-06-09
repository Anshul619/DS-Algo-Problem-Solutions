package main

import (
	"testing"
)

func TestOnlineStockSpan1(t *testing.T) {
	ss := Constructor()

	stocks := []struct {
		stock    int
		expected int
	}{
		{100, 1},
		{80, 1},
		{60, 1},
		{70, 2},
		{60, 1},
		{75, 4},
		{85, 6},
	}

	for _, v := range stocks {
		if ss.Next(v.stock) != v.expected {
			t.Errorf("failed %v", v.stock)
		}
	}
}

func TestOnlineStockSpan2(t *testing.T) {
	ss := Constructor()

	stocks := []struct {
		stock    int
		expected int
	}{
		{31, 1},
		{41, 2},
		{48, 3},
		{59, 4},
		{79, 5},
	}

	for _, v := range stocks {
		if ss.Next(v.stock) != v.expected {
			t.Errorf("failed %v", v.stock)
		}
	}
}
