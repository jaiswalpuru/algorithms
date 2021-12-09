/*
Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

MyHashMap() initializes the object with an empty map.
void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map,
update the corresponding value.
int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
*/

package main

import "fmt"

type HashMap struct {
	arr []int
}

func Constructor() HashMap {
	mp := make([]int, 1000001)
	for i := 0; i < 1000001; i++ {
		mp[i] = -1
	}
	hp := HashMap{
		arr: mp,
	}
	return hp
}

func (this *HashMap) Put(key, val int) { this.arr[key] = val }
func (this *HashMap) Get(key int) int  { return this.arr[key] }
func (this *HashMap) Remove(key int)   { this.arr[key] = -1 }

func main() {
	obj := Constructor()
	obj.Put(1, 2)
	fmt.Println(obj.Get(1))
	obj.Remove(1)
	fmt.Println(obj.Get(1))
}
