package main

/*
- Leetcode - https://leetcode.com/problems/online-stock-span/description/
- Space - O(n)
- Time - O(n)
*/

type Stock struct {
	val int
	day int
}

type stackStocks []*Stock

func (s *stackStocks) push(st *Stock) {
	(*s) = append((*s), st)
}

func (s *stackStocks) pop() *Stock {
	temp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return temp
}

func (s *stackStocks) peek() *Stock {
	return (*s)[len(*s)-1]
}

func (s stackStocks) isEmpty() bool {
	return len(s) == 0
}

type StockSpanner struct {
	stack      *stackStocks
	currentDay int
}

func Constructor() StockSpanner {
	s := new(stackStocks)

	return StockSpanner{
		stack:      s,
		currentDay: 0,
	}
}

func (this *StockSpanner) Next(price int) int {

	span := 1

	this.currentDay++

	for !this.stack.isEmpty() && this.stack.peek().val <= price {
		this.stack.pop()
	}

	if !this.stack.isEmpty() {
		span = this.currentDay - this.stack.peek().day
	} else {
		span = this.currentDay
	}

	this.stack.push(&Stock{
		val: price,
		day: this.currentDay,
	})

	return span
}

// func main() {
// 	// stockSpanner := Constructor()
// 	// log.Println(stockSpanner.Next(100))
// 	// log.Println(stockSpanner.Next(80))
// 	// log.Println(stockSpanner.Next(60))
// 	// log.Println(stockSpanner.Next(70))
// 	// log.Println(stockSpanner.Next(60))
// 	// log.Println(stockSpanner.Next(75))
// 	// log.Println(stockSpanner.Next(85))

// 	// stockSpanner := Constructor()
// 	// log.Println(stockSpanner.Next(31))
// 	// log.Println(stockSpanner.Next(41))
// 	// log.Println(stockSpanner.Next(48))
// 	// log.Println(stockSpanner.Next(59))
// 	// log.Println(stockSpanner.Next(79))

// 	stockSpanner := Constructor()
// 	log.Println(stockSpanner.Next(28))
// 	log.Println(stockSpanner.Next(14))
// 	log.Println(stockSpanner.Next(28))
// 	log.Println(stockSpanner.Next(35))
// 	log.Println(stockSpanner.Next(46))
// 	log.Println(stockSpanner.Next(53))
// 	log.Println(stockSpanner.Next(66))
// 	log.Println(stockSpanner.Next(80))
// 	log.Println(stockSpanner.Next(87))
// 	log.Println(stockSpanner.Next(88))

// }
