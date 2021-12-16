/*
Author : Puru Jaiswal
References : CLRS
*/

package queue

import "fmt"

type Item interface{}

type Queue struct {
	data        []Item
	front, rear int
}

func New() *Queue { return &Queue{front: -1, rear: -1} }

func (q *Queue) Enqueue(val Item) {
	if q.rear == -1 {
		q.data = make([]Item, 1)
		q.front = 0
		q.rear = 0
		q.data[q.rear] = val
	} else {
		q.data = append(q.data, val)
		q.rear++
	}
}

func (q *Queue) Dequeue() Item {
	if q.front == -1 {
		fmt.Println("Queue Empty")
		return nil
	}
	val := q.data[q.front]
	q.front++
	if q.front > q.rear {
		q.front = -1
		q.rear = -1
		q.data = []Item{}
	}
	return val
}
