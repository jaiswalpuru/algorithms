package main

import "fmt"

type Item struct {
	key   int
	value int
}

type HashMap struct {
	size int
	db   [][]Item
}

func (h *HashMap) _constructor(size int) {
	h.size = size
	h.db = make([][]Item, size)
	for i := 0; i < h.size; i++ {
		h.db[i] = make([]Item, 0)
	}
}

func (h *HashMap) _hash_function(key int) int { return key % h.size }

func (h *HashMap) set(key, val int) {
	hash_ind := h._hash_function(key)
	for i := 0; i < len(h.db[hash_ind]); i++ {
		if h.db[hash_ind][i].value == key {
			h.db[hash_ind][i].value = val
			return
		}
	}
	h.db[hash_ind] = append(h.db[hash_ind], Item{key: key, value: val})
}

func (h *HashMap) get(key int) int {
	hash_ind := h._hash_function(key)
	for i := 0; i < len(h.db[hash_ind]); i++ {
		if h.db[hash_ind][i].key == key {
			return h.db[hash_ind][i].value
		}
	}
	fmt.Println("Key not found")
	return -1
}

func (h *HashMap) remove(key int) {
	hash_index := h._hash_function(key)
	for i := 0; i < len(h.db[hash_index]); i++ {
		if h.db[hash_index][i].key == key {
			t1 := h.db[hash_index][:i]
			t2 := h.db[hash_index][i+1:]
			res := []Item{}
			res = append(res, t1...)
			res = append(res, t2...)
			h.db[hash_index] = res
			return
		}
	}
	fmt.Println("Key not found")
}
