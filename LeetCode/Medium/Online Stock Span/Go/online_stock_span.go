package main

type StockSpanner struct{ stack []Pair }

type Pair struct{ val, cnt int }

func Constructor() StockSpanner {
	return StockSpanner{stack: []Pair{}}
}

func (this *StockSpanner) Next(price int) int {
	cnt := 1
	for len(this.stack) > 0 && this.stack[len(this.stack)-1].val <= price {
		cnt += this.stack[len(this.stack)-1].cnt
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, Pair{price, cnt})
	return cnt
}
