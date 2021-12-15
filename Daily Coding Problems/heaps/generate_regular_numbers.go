/*
	A regular number in mathematics id defined as one which evenly divides some power of 60.
	Equivalently, we can say that a regular number is one whose only prime divisors are 2,3 and 5.
	Given a number n, write a program that generates, in order, the first n regular numbers.
*/

package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

/*
	Naive soln would be to first generate all powers of 2,3 and 5 up to some stopping point, and then find every
	product we can obtain from multiplying one power from each group. Then we can sort these products and take
	the first n to find our soln.
*/

func naive_soln(n int) []int {
	arr_two := []int{}
	arr_three := []int{}
	arr_five := []int{}

	for i := 0; i < n; i++ {
		arr_two = append(arr_two, int(math.Pow(2, float64(i))))
		arr_three = append(arr_three, int(math.Pow(3, float64(i))))
		arr_five = append(arr_five, int(math.Pow(5, float64(i))))
	}

	mp := make(map[int]bool)
	soln := []int{}

	for _, i := range arr_two {
		for _, j := range arr_three {
			for _, k := range arr_five {
				if _, ok := mp[i*j*k]; !ok {
					soln = append(soln, i*j*k)
					mp[i*j*k] = true
				}
			}
		}
	}
	sort.Ints(soln)
	return soln[:n]
}

/*
	Efficient approach : For any regular number x we can generate three additional regular numbers by calculating
	2x, 3x, and 5x.
*/

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinHeap) Push(val interface{}) { *h = append(*h, val.(int)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func regular_numbers(n int) {
	soln := make(MinHeap, 0)
	last, cnt := 0, 0

	heap.Init(&soln)
	heap.Push(&soln, 1)
	list := []int{}
	for cnt < n {
		x := heap.Pop(&soln).(int)
		if x > last {
			list = append(list, x)
			last = x
			cnt += 1
			heap.Push(&soln, 2*x)
			heap.Push(&soln, 3*x)
			heap.Push(&soln, 5*x)
		}
	}

	fmt.Println(list)
}

func main() {
	regular_numbers(10)
}
