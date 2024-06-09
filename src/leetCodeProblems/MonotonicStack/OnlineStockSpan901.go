package main

/*
- Leetcode - https://leetcode.com/problems/online-stock-span/description/
- Space - O(n)
- Time - O(n)
*/

type Stock struct {
	price int
	span  int
}

type stackStocks []Stock

func (s *stackStocks) push(st Stock) {
	(*s) = append((*s), st)
}

func (s *stackStocks) pop() Stock {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *stackStocks) peek() Stock {
	return (*s)[len(*s)-1]
}

func (s stackStocks) isEmpty() bool {
	return len(s) == 0
}

type StockSpanner struct {
	stack *stackStocks
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: new(stackStocks),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1

	for !this.stack.isEmpty() &&
		this.stack.peek().price <= price {
		span += this.stack.pop().span
	}

	this.stack.push(Stock{
		price: price,
		span:  span,
	})
	return span
}
