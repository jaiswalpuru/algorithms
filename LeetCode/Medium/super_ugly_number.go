/*
A super ugly number is a positive integer whose prime factors are in the array primes.

Given an integer n and an array of integers primes, return the nth super ugly number.

The nth super ugly number is guaranteed to fit in a 32-bit signed integer.
*/

package main

import (
	"container/heap"
	"fmt"
	"math"
)

type MinHeap []int

func (pq MinHeap) Len() int { return len(pq) }
func (pq MinHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}
func (pq MinHeap) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *MinHeap) Pop() interface{} {
	x := (*pq)[(*pq).Len()-1]
	*pq = (*pq)[:(*pq).Len()-1]
	return x
}

func (pq *MinHeap) Push(x interface{}) { *pq = append(*pq, x.(int)) }

//this wont work cannot allocate space in heap
func super_ugly_number(n int, prime []int) int {
	mh := MinHeap{1}
	heap.Init(&mh)

	visited := map[int]bool{1: true}
	m := len(prime)
	var res int
	for n > 0 {
		res = heap.Pop(&mh).(int)
		for i := 0; i < m; i++ {
			temp := res * prime[i]
			if !visited[temp] {
				visited[temp] = true
				heap.Push(&mh, temp)
			}
		}
		n--
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func super_ugly_number_dp(n int, prime []int) int {
	m := len(prime)
	ugly := make([]int, n)
	index := make([]int, m)

	ugly[0] = 1
	for i := 1; i < n; i++ {
		ugly[i] = int(math.MaxInt64)
		for j := 0; j < m; j++ {
			ugly[i] = min(ugly[i], prime[j]*ugly[index[j]])
		}

		for j := 0; j < m; j++ {
			for prime[j]*ugly[index[j]] <= ugly[i] {
				index[j]++
			}
		}
	}
	return ugly[n-1]
}

func main() {
	fmt.Println(super_ugly_number_dp(12, []int{2, 7, 13, 19}))
}
