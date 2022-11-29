package main

import "math/rand"

type RandomizedSet struct {
	hash_map map[int]int
	arr      []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{hash_map: make(map[int]int), arr: make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash_map[val]; ok {
		return false
	}
	this.hash_map[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hash_map[val]; ok {
		last_val := this.arr[len(this.arr)-1]
		ind := this.hash_map[val]
		this.arr[ind] = last_val
		this.hash_map[last_val] = ind
		this.arr = this.arr[:len(this.arr)-1]
		delete(this.hash_map, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
