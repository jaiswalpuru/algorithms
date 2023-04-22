package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Pair struct{ enqueTime, processTime, index int }

// initialize heap interface
type M []Pair

func (m M) Len() int      { return len(m) }
func (m M) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m M) Less(i, j int) bool {
	if m[i].processTime == m[j].processTime {
		return m[i].index < m[j].index
	}
	return m[i].processTime < m[j].processTime
}
func (m *M) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func getOrder(tasks [][]int) []int {
	size := len(tasks)
	tasksMod := make([]Pair, size)
	for i := 0; i < size; i++ {
		tasksMod[i].enqueTime = tasks[i][0]
		tasksMod[i].processTime = tasks[i][1]
		tasksMod[i].index = i
	}

	// initially sort by enqueue time
	sort.Slice(tasksMod, func(i, j int) bool {
		if tasksMod[i].enqueTime == tasksMod[j].enqueTime {
			return tasksMod[i].processTime < tasksMod[j].processTime
		}
		return tasksMod[i].enqueTime < tasksMod[j].enqueTime
	})

	//create a min heap for
	minHeap := &M{}
	heap.Push(minHeap, tasksMod[0])
	i := 1 //index for tasksMod
	startTime := tasksMod[0].enqueTime

	ordering := make([]int, size)
	k := 0 //index for ordering

	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(Pair)
		ordering[k] = top.index
		k++
		startTime += top.processTime

		//push all the element to heap whose enque Time is less than start time
		for i < size && tasksMod[i].enqueTime <= startTime {
			heap.Push(minHeap, tasksMod[i])
			i++
		}

		//if the length of the heap is zero and still processes are left to be procesed
		if minHeap.Len() == 0 && i < size {
			heap.Push(minHeap, tasksMod[i])
			startTime += tasksMod[i].enqueTime
			i++
		}
	}
	return ordering
}

func main() {
	fmt.Println(getOrder([][]int{
		{1, 2}, {2, 4}, {3, 2}, {4, 1},
	}))
}
