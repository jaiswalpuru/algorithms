package main

import "container/heap"

type Node struct {
	val int
	index int
	cnt int
}

type I []Node
func (mh I) Len() int { return len(mh)}
func (mh I) Less(i,j int) bool {
	if mh[i].cnt != mh[j].cnt {
		return mh[i].cnt > mh[j].cnt
	}else{
		return mh[i].index > mh[j].index
	}
}
func (mh I) Swap(i,j int) { mh[i],mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(Node)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

type FreqStack struct {
	f_stack I
	ind int
	hash_map map[int]int
}

func Constructor() FreqStack { 
	var max_heap FreqStack
	max_heap.ind = 0
	max_heap.hash_map = make(map[int]int)
	heap.Init(&max_heap.f_stack)
	return max_heap
}


func (this *FreqStack) Push(val int)  {
	this.ind+=1
	this.hash_map[val]++
    heap.Push(&this.f_stack, Node{val, this.ind, this.hash_map[val]})
}


func (this *FreqStack) Pop() int {
    val := heap.Pop(&this.f_stack).(Node)
    this.hash_map[val.val]--
    return val.val
}


func main() {}