package main

import "container/list"

type Element struct {
	key,
	val int
}

type LRUCache struct {
	capacity int
	l        *list.List
	hash_map map[int]*list.Element
}

func Constructor(cap int) LRUCache {
	return LRUCache{capacity: cap, l: list.New(), hash_map: make(map[int]*list.Element)}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hash_map[key]; ok {
		this.l.MoveToFront(v)
		return v.Value.(*Element).val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.hash_map[key]
	if ok {
		v.Value.(*Element).val = value
		this.l.MoveToFront(v)
		return
	}

	if this.l.Len() == this.capacity {
		//removing the rear element i.e the oldest element
		ele := this.l.Back()
		this.l.Remove(ele)
		delete(this.hash_map, ele.Value.(*Element).key)
	}

	item := &Element{key: key, val: value}
	element := this.l.PushFront(item)
	this.hash_map[key] = element
}

func main() {

}
