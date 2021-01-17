package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

var a, b []int

type pair struct {
	p, q int
}

func (x pair) value() int {
	return a[x.p] + b[x.q]
}

type pairHeap []pair

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool { return h[i].value() < h[j].value() }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }

func (h *pairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func fillArray(a []int, n int) {
	for i := 0; i < n; i++ {
		if i == n-1 {
			scanf("%d\n", &a[i])
			break
		}
		scanf("%d ", &a[i])
	}
}

func main() {
	defer writer.Flush()

	var (
		t, k, Q int
	)

	scanf("%d\n", &t)

	for ; t > 0; t-- {
		scanf("%d %d\n", &k, &Q)

		a = make([]int, k)
		b = make([]int, k)

		fillArray(a, k) // take inputs for a
		fillArray(b, k) //take inputs for b

		sort.Ints(a)
		sort.Ints(b)

		queries := make([]int, Q)
		maxQuery := -1
		for i := 0; i < Q; i++ {
			if i == Q-1 {
				scanf("%d\n", &queries[i])
				if maxQuery < queries[i] {
					maxQuery = queries[i]
				}
				break
			}
			scanf("%d ", &queries[i])
			if maxQuery < queries[i] {
				maxQuery = queries[i]
			}
		}
		h := make(pairHeap, k)
		for i := 0; i < k; i++ {
			h[i].p = i
			h[i].q = 0
		}
		heap.Init(&h)
		generated := make([]int, 0, maxQuery)
		for i := 0; i < maxQuery; i++ {
			generated = append(generated, h[0].value())
			h[0].q++
			if h[0].q < k {
				heap.Fix(&h, 0)
			} else {
				heap.Pop(&h)
			}
		}

		for _, q := range queries {
			printf("%d\n", generated[q-1])
		}
	}
}
