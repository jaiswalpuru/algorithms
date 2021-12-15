/*
Author : Puru Jaiswal
References : CLRS
*/

package heaps

import (
	"fmt"
	"math"
)

//interface for min/max heap check
type Item interface{ Less(item Item) bool }

//heap structure
type Heap struct {
	size int
	data []Item
}

//return the parent of the current node
func parent(i int) int { return int(math.Floor(float64(i-1) / 2)) }

//left child
func left_child(parent int) int { return (2 * parent) + 1 }

//right child
func right_child(parent int) int { return (2 * parent) + 2 }

func New() *Heap { return &Heap{} }

//as this is a min heap so the smallest element will always be at the root in contrast to max heap the largest element
//is at the root
func get_minimum(h *Heap) Item {
	if h.size == 0 {
		fmt.Println("Heap is empty")
		return nil
	}
	return h.data[0]
}

//if the Less function is implemented for max heap then
func get_maximum(h *Heap) Item {
	if h.size == 0 {
		fmt.Println("Heap is empty")
		return nil
	}
	return h.data[0]
}

func swap(h *Heap, i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }

//percolate up moves the element from bottom to top
func (h *Heap) percolate_up() {

	idx := h.size - 1
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

//percolate down moves the element from top to bottom which will be called while deletion
/*
Deletion process
1. Swap the last element and the root element
2. Remove the last element from heap
3. Recursively call percolate down on root to place it to its correct location
*/
func (h *Heap) percolate_down(i int) {

	p := i

	for {

		left := left_child(p)
		right := right_child(p)
		s := p

		if left < h.size && h.data[left].Less(h.data[s]) {
			s = left
		}

		if right < h.size && h.data[right].Less(h.data[s]) {
			s = right
		}

		if s == p {
			break
		}

		swap(h, s, p)
		p = s

	}
}

//Extract return the root element from the node and performs heapify
func (h *Heap) Extract() Item {

	if h.size == 0 {
		fmt.Println("Empty Heap")
		return nil
	}

	data := h.data[0]
	//swap the root with last element
	h.data[0] = h.data[h.size-1]
	//remove the last element
	h.data = h.data[:h.size-1]
	h.size--

	//call percolate_down on the root node
	if h.size > 0 {
		h.percolate_down(0)
	} else {
		h.data = nil
	}

	return data
}

//Insert
func (h *Heap) Insert(data Item) {

	if h.size == 0 {
		h.data = make([]Item, 1)
		h.data[0] = data
	} else {
		h.data = append(h.data, data)
	}

	h.size++
	h.percolate_up()

}

//Heapify
func Heapify(items []Item) *Heap {

	h := New()
	h.size = len(items)
	h.data = make([]Item, h.size)

	copy(h.data, items)

	i := int(h.size / 2)

	for i >= 0 {
		h.percolate_down(i)
		i--
	}

	return h
}
