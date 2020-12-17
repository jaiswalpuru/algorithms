package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

//IntHeap represents a heap with elements of datatype as integers
type IntHeap []int

//Len - return the len of the heap
func (h IntHeap) Len() int { return len(h) }

//Less - returns whether x < y (true or false)
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

//Swap - is used to swap the respective values
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

//Push - appends an elements onto the heap
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

//Pop - removes the last element from the heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		var n, z int
		scanf("%d %d\n", &n, &z)

		arr := make(IntHeap, n)

		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
				break
			}
			scanf("%d ", &arr[i])
		}

		//For priority queue
		heap.Init(&arr)
		ans := 0
		for z > 0 && len(arr) > 0 {
			ans++
			z -= arr[0]
			arr[0] /= 2
			if arr[0] == 0 {
				heap.Pop(&arr)
			} else {
				heap.Fix(&arr, 0)
			}
		}

		if z > 0 {
			printf("%s\n", "Evacuate")
		} else {
			printf("%d\n", ans)
		}

	}
}
