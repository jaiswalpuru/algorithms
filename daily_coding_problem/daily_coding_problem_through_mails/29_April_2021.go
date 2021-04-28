// You run an e-commerce website and want to record the last N order ids in a log.
// Implement a data structure to accomplish this, with the following API:

// record(order_id): adds the order_id to the log
// get_last(i): gets the ith last element from the log. i is guaranteed to be smaller than or equal to N.

package main

import "fmt"

type OrderID int

type Log struct {
	size   int
	logArr []OrderID
	index  int
}

func NewLog(size int) *Log {
	return &Log{size: size, logArr: make([]OrderID, size), index: 0}
}

func (l *Log) record(orderID OrderID) {
	l.logArr[l.index] = orderID
	l.index = (l.index + 1) % l.size
}

func (l *Log) get_last(ind int) OrderID {
	if ind > l.size {
		ind = ind % l.size
	}
	return l.logArr[(l.index-ind+l.size)%l.size]
}

func main() {
	nl := NewLog(3)
	nl.record(25)
	nl.record(30)
	nl.record(40)
	fmt.Println(nl.get_last(5))
	fmt.Println(nl.get_last(0))
}
