package main

import "testing"

func TestTimeNeededToBuyTicket(t *testing.T) {
	tests := []struct {
		tickets  []int
		k        int
		expected int
	}{
		{[]int{2, 3, 2}, 2, 6},
		{[]int{5, 1, 1, 1}, 0, 8},
	}
	for i, v := range tests {
		if timeRequiredToBuy(v.tickets, v.k) != v.expected {
			t.Errorf("failed %v", i)
		}
	}
}
