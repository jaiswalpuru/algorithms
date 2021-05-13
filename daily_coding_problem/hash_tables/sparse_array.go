/*
	You have a large array, most of whose elements are zero.
	Create a more space efficient data structure, SparseArray that implements the following interface
	• init(arr, size) : initialize with the original large array and size.
	• set(i, val) : update index at i to be val
	• get(i) : get the value at index i
*/

package main

import "fmt"

type SparseArray struct {
	size int
	mp   map[int]int
}

//arr -> here it is large array
func init_(arr []int, size int) *SparseArray {
	sp := &SparseArray{size: size}
	sp.mp = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			sp.mp[i] = arr[i]
		}
	}
	return sp
}

func (sp *SparseArray) _check_bound(i int) {
	if i < 0 || i > sp.size {
		panic("error : Index out of bounds")
	}
}

func (sp *SparseArray) set(i, val int) {
	sp._check_bound(i)
	if val != 0 {
		sp.mp[i] = val
	} else if _, ok := sp.mp[i]; ok {
		delete(sp.mp, i)
	}
}

func (sp *SparseArray) get(i int) int {
	sp._check_bound(i)
	if _, ok := sp.mp[i]; ok {
		return sp.mp[i]
	} else {
		return 0
	}
}

func main() {
	sp := init_([]int{1, 2, 3, 4, 5, 6}, 6)
	fmt.Println(sp.get(4))
	fmt.Println(sp.get(8))
}
