package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Int - used for alernative type
type Int int

//Less - return
func (a Int) Less(b Item) bool {
	val, ok := b.(Int)
	return ok && a <= val
}

func verifyHeap(h *Heap) bool {
	queue := make([]Int, 1)
	queue[0] = 0
	for {
		p := queue[0]
		queue = queue[1:]
		l := leftChild(int(p))
		r := rightChild(int(p))
		if l < h.size {
			if !h.data[p].Less(h.data[l]) {
				return false
			}
			queue = append(queue, Int(l))
		}
		if r < h.size {
			if !h.data[p].Less(h.data[r]) {
				return false
			}
			queue = append(queue, Int(r))
		}
	}
	return true
}

func verifyStrictlyIncreasing(h *Heap) (bool, []Item) {
	prev, _ := h.Extract()
	order := []Item{prev}
	for h.size > 0 {
		cur, _ := h.Extract()
		order = append(order, cur)
		if cur.Less(prev) {
			return false, order
		}
		prev = cur
		order = append(order, prev)
	}
	return true, order
}

func randomPerm(n int) []Item {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	ints := r.Perm(n)
	items := make([]Item, n)
	for idx, item := range ints {
		items[idx] = Int(item)
	}
	return items
}

func main() {
	items := randomPerm(20)
	hp := New()
	for _, item := range items {
		fmt.Println("Inserting an element into Heap", hp.data)
		hp.Insert(item)
	}
	if !verifyHeap(hp) {
		fmt.Println("Invalid heap : ", hp.data)
		return
	}
	if ok, order := verifyStrictlyIncreasing(hp); !ok {
		fmt.Println("Invalid Heap extraction order : ", order)
	}

	items2 := randomPerm(30)
	for i := 0; i < len(items2); i++ {
		fmt.Println(items2[i].(Int), " ")
	}
	items2 = HeapSort(items2)
	fmt.Println()
	for i := 0; i < len(items2); i++ {
		fmt.Println(items2[i].(Int), " ")
	}

}
