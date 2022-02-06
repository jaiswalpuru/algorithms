package main

import (
	"strings"
)

type Bitset struct {
	size              int
	arr               []string
	flip_arr          []string
	cnt_one, cnt_zero int
}

func Constructor(size int) Bitset {
	arr1 := make([]string, size)
	arr2 := make([]string, size)
	for i := 0; i < size; i++ {
		arr1[i] = "0"
		arr2[i] = "1"
	}
	b := Bitset{size: size, arr: arr1, flip_arr: arr2, cnt_zero: size, cnt_one: 0}
	return b
}

func (this *Bitset) Fix(idx int) {
	if this.arr[idx] != "1" {
		this.cnt_zero--
		this.cnt_one++
	}
	this.arr[idx] = "1"
	this.flip_arr[idx] = "0"
}

func (this *Bitset) Unfix(idx int) {
	if this.arr[idx] != "0" {
		this.cnt_zero++
		this.cnt_one--
	}
	this.arr[idx] = "0"
	this.flip_arr[idx] = "1"
}

func (this *Bitset) Flip() {
	this.arr, this.flip_arr = this.flip_arr, this.arr
	this.cnt_one, this.cnt_zero = this.cnt_zero, this.cnt_one
}

func (this *Bitset) All() bool {
	return this.cnt_one == this.size
}

func (this *Bitset) One() bool {
	return this.cnt_one > 0
}

func (this *Bitset) Count() int {
	return this.cnt_one
}

func (this *Bitset) ToString() string {
	return strings.Join(this.arr, "")
}

func main() {
	obj := Constructor(5)
	obj.Fix(3)
	obj.Unfix(3)
	obj.Flip()
}
