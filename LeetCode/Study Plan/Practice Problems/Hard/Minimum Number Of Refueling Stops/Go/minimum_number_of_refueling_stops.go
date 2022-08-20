package main

import (
	"container/heap"
	"fmt"
)

type MH []int

func (m MH) Len() int              { return len(m) }
func (m MH) Less(i, j int) bool    { return m[i] > m[j] }
func (m MH) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *MH) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *MH) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func minimum_number_of_refueling_stops(target, start_fuel int, stations [][]int) int {
	m := &MH{}
	ans, prev := 0, 0
	n := len(stations)
	for i := 0; i < n; i++ {
		location := stations[i][0]
		capacity := stations[i][1]
		start_fuel -= location - prev
		for m.Len() > 0 && start_fuel < 0 {
			start_fuel += heap.Pop(m).(int)
			ans++
		}
		if start_fuel < 0 {
			return -1
		}
		heap.Push(m, capacity)
		prev = location
	}

	start_fuel -= target - prev
	for m.Len() > 0 && start_fuel < 0 {
		start_fuel += heap.Pop(m).(int)
		ans++
	}
	if start_fuel < 0 {
		return -1
	}

	return ans
}

func main() {
	fmt.Println(minimum_number_of_refueling_stops(100, 10, [][]int{
		{10, 60}, {20, 30}, {30, 30}, {60, 40},
	}))
}
