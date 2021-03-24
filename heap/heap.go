package main

import (
	"fmt"
	"math"
)

//Item - defines the interface of an element to be held by a Heap instance.
type Item interface {
	Less(item Item) bool
}

//Heap - binary heap with support for the min heap operations
type Heap struct {
	size int
	data []Item
}

//New - returns a pointer to an empty min -heap
func New() *Heap { return &Heap{} }

//returns the parent of a node assume stored in array then parent is present at position (i-1)/2
func parent(i int) int { return int(math.Floor(float64(i-1) / 2.0)) }

//returns the left child for a node
func leftChild(parent int) int { return (2 * parent) + 1 }

//returns the right child for a node
func rightChild(parent int) int { return (2 * parent) + 2 }

//Get the minimum element if heap is a min heap then element is present at the 0th index of the array
func getMinimum(h *Heap) (Item, error) {
	if h.size == 0 {
		return nil, fmt.Errorf("Unable to get the element from an empty heap")
	}
	return h.data[0], nil
}

//swapping elements
func swap(h *Heap, i, j int) {
	tmp := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = tmp
}

//percolateUp is the process in which element is moved from bottom to top which occurs in the case of insertion
//Percolate Up is called when an element is inserted in heap, as the element is added at the end of the array
func (h *Heap) percolateUp() {
	idx := h.size - 1 //getting the index of last element
	if idx <= 0 {
		return
	}

	for {
		p := parent(idx)
		if p < 0 || h.data[p].Less(h.data[idx]) {
			break
		}
		swap(h, p, idx)
		idx = p
	}
}

//percolateDown is the process in which element is moved from top to bottom which occurs in case of deletion
//deletion process
// 1 - Swap the last element and the top element
// 2 - then move the element to its respecticve position so that it satisfies the heap property
// 3 - delete the last element
func (h *Heap) percolateDown(i int) {
	p := i
	for {
		l := leftChild(p)
		r := rightChild(p)
		s := p
		if l < h.size && h.data[l].Less(h.data[s]) {
			s = l
		}
		if r < h.size && h.data[r].Less(h.data[s]) {
			s = r
		}
		if s == p {
			break
		}
		swap(h, p, s)
		p = s
	}
}

//Extract - removes and return the item at the top of the heap maintaining the heap properties
func (h *Heap) Extract() (Item, error) {
	n := h.size
	if n == 0 {
		return nil, fmt.Errorf("Unable to extract from an empty heap")
	}
	m := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	h.size--
	if h.size > 0 {
		h.percolateDown(0)
	} else {
		h.data = nil
	}
	return m, nil
}

//Insert - inserts item into heap maintaining the min heap
func (h *Heap) Insert(item Item) {
	if h.size == 0 {
		h.data = make([]Item, 1)
		h.data[0] = item
	} else {
		h.data = append(h.data, item)
	}
	h.size++
	h.percolateUp() //bottom to top
}

//Heapify - return a pointer to a min heap composed of elements of items
func Heapify(items []Item) *Heap {
	h := New()

	n := len(items)

	h.data = make([]Item, n)
	copy(h.data, items)
	h.size = n

	i := int(n / 2)

	for i >= 0 {
		h.percolateDown(i)
		i--
	}
	return h
}
