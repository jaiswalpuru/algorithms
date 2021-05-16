/*
Compute the running median of a sequence of numbers. That is, given a stream of numbers,
print out the median of the list so far on each new element.

Recall that the median of an even-numbered list is the average of the two middle numbers.

For example, given the sequence [2, 1, 5, 7, 2, 0, 5], your algorithm should print out:

2
1.5
2
3.5
2
2
2

Brute force approach is to just apply insertion sort on the incomming elements and median can be found in O(1).
TC : O(n^2)

Efficient Soln : Use two heaps maxHeap and minHeap. maxHeap will have first half of the elements and minHeap
will have second half of elements.

*/

package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int
type MinHeap []int

//MaxHeap
func (mh MaxHeap) Len() int           { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool { return mh[i] > mh[j] }
func (mh MaxHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MaxHeap) Push(x interface{}) { *mh = append(*mh, x.(int)) }

func (mh *MaxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	val := old[n-1]
	*mh = old[:n-1]
	return val
}

//MinHeap
func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Push(x interface{}) { *mh = append(*mh, x.(int)) }

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	val := old[n-1]
	*mh = old[:n-1]
	return val
}

func main() {
	arr := []int{2, 1, 5, 7, 2, 0, 5}
	maxHeap := make(MaxHeap, 0)
	minHeap := make(MinHeap, 0)

	heap.Init(&maxHeap)
	heap.Init(&minHeap)

	for i := 0; i < len(arr); i++ {
		lSize := maxHeap.Len()
		rSize := minHeap.Len()

		if lSize == 0 {
			heap.Push(&maxHeap, arr[i])
		} else if lSize == rSize {
			if minHeap[0] > arr[i] {
				heap.Push(&maxHeap, arr[i])
			} else {
				val := heap.Pop(&minHeap)
				heap.Push(&maxHeap, val)
				heap.Push(&minHeap, arr[i])
			}
		} else {
			if rSize == 0 {
				if arr[i] > maxHeap[0] {
					heap.Push(&minHeap, arr[i])
				} else {
					val := heap.Pop(&maxHeap)
					heap.Push(&minHeap, val)
					heap.Push(&maxHeap, arr[i])
				}
			} else if arr[i] >= minHeap[0] {
				heap.Push(&minHeap, arr[i])
			} else {
				if maxHeap[0] > arr[i] {
					val := heap.Pop(&maxHeap)
					heap.Push(&minHeap, val)
					heap.Push(&maxHeap, arr[i])
				} else {
					heap.Push(&minHeap, arr[i])
				}
			}
		}
		if (i+1)%2 == 0 {
			fmt.Printf("i = %d, median = %.1f\n", i+1, float64(maxHeap[0]+minHeap[0])/float64(2))
		} else {
			fmt.Printf("i = %d, median = %d\n", i+1, maxHeap[0])
		}
	}

}
