package main

type node struct {
	price int
	span  int
}

// Medium: OnlineStockSpan
// Solution for online-stock-span (medium)
type StockSpanner struct {
	stack []node
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (stockSpanner *StockSpanner) Next(price int) int {
	current := node{
		price,
		1,
	}

	for len(stockSpanner.stack) > 0 && stockSpanner.stack[len(stockSpanner.stack)-1].price <= price {
		pastNode := stockSpanner.stack[len(stockSpanner.stack)-1]
		current.span += pastNode.span
		stockSpanner.stack = stockSpanner.stack[:len(stockSpanner.stack)-1]
	}

	stockSpanner.stack = append(stockSpanner.stack, current)

	return current.span
}
