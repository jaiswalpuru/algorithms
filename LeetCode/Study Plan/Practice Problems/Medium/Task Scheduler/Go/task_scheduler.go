package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//-------------------------using heap-----------------------
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

type Pair struct{ val, time int }

func task_scheduler_heap(tasks []byte, n int) int {
	mp := make(map[byte]int)
	for i := 0; i < len(tasks); i++ {
		mp[tasks[i]]++
	}
	max_heap := &M{}
	for _, v := range mp {
		heap.Push(max_heap, v)
	}
	q := []Pair{}
	time := 0
	for max_heap.Len() > 0 || len(q) > 0 {
		time += 1
		if max_heap.Len() > 0 {
			cnt := heap.Pop(max_heap).(int) - 1
			if cnt != 0 {
				q = append(q, Pair{cnt, time + n})
			}
		}
		if len(q) > 0 {
			curr := q[0]
			if curr.time == time {
				q = q[1:]
				heap.Push(max_heap, curr.val)
			}
		}
	}
	return time
}

//-------------------------using heap-----------------------

func task_scheduler(tasks []byte, n int) int {
	freq := make([]int, 26)
	for i := 0; i < len(tasks); i++ {
		freq[int(tasks[i])-int('A')]++
	}
	sort.Ints(freq)
	max_freq := freq[25]
	idle_time := (max_freq - 1) * n
	for i := len(freq) - 2; i >= 0; i-- {
		idle_time -= min(max_freq-1, freq[i])
	}
	idle_time = max(idle_time, 0)
	return idle_time + len(tasks)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(task_scheduler([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
}
