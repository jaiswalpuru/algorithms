package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()
	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(14)
	//q.data[0] = 14 //Introduing error for test purpose
	val := q.Dequeue()
	if val != 12 {
		t.Error("Error while dequeing")
	}
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}
