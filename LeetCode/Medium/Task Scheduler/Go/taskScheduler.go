package main

import (
	"container/heap"
	"fmt"
)

type M []int

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i] > m[j] }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

type Pair struct {
	taskCnt        int
	coolDownPeriod int
}

func leastInterval(tasks []byte, n int) int {
	size := len(tasks)
	tasksCount := make(map[byte]int)
	for i := 0; i < size; i++ {
		tasksCount[tasks[i]]++
	}

	maxHeap := &M{}
	for _, v := range tasksCount {
		heap.Push(maxHeap, v)
	}

	time := 0
	q := []Pair{}
	for len(q) > 0 || maxHeap.Len() > 0 {
		time += 1

		if maxHeap.Len() > 0 {
			top := heap.Pop(maxHeap).(int) - 1
			if top != 0 { //still the task has other parts
				q = append(q, Pair{top, time + n})
			}
		}

		if len(q) > 0 {
			front := q[0]
			if time == front.coolDownPeriod {
				q = q[1:]
				heap.Push(maxHeap, front.taskCnt)
			}
		}
	}

	return time
}

func main() {
	fmt.Println(leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
