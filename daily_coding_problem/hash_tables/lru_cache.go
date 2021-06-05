/*
Implement a LRU cache
Least recently used cache. The cache should be able to initialized with cache size n, and provide the following details.
• set (key, value): set key to value. If there are already n items in the cache and we are adding a new item,
  also remove the least recently used item.
• get(key) : get the value of key. If no such key exists, return null.

Create a map of key type and node of the linked list
Add the node at the end of the list
Remove from beginning

Set :
While setting the cache check if it present or not.
If present in map delete the value.
Add in list
Check if length of cache doesn't exceed the cache size
Add in the map
if length exceeds the cache size remove the node from front of the list as it is least used node

Get:
Check if it is present in map, if it is not present return -1
If present, remove from list then again add at the end of list.
return the value of the node
*/

package main

import "fmt"

type DLL struct {
	key  int
	data int
	prev *DLL
	next *DLL
}

func NewLL(key, data int) *DLL { return &DLL{key: key, data: data, prev: nil, next: nil} }

var (
	//dummy nodes
	head = NewLL(-1, -1)
	tail = NewLL(-1, -1)
	mp   = make(map[int]*DLL)
	n    = 5 //cache of size 5
)

func getHead() *DLL {
	return head.next
}

func getTail() *DLL {
	return tail.prev
}

//adds at the end of the dll (exactly before the 't' a dummy node)
func add(node *DLL) {
	p := tail.prev
	p.next = node
	node.prev = p
	node.next = tail
	tail.prev = node
}

func remove(node *DLL) {
	p := node.prev
	nxt := node.next
	p.next = nxt
	nxt.prev = p
}

func set(key, val int) {
	if _, ok := mp[key]; ok {
		delete(mp, key)
	}
	node := NewLL(key, val)
	add(node)
	mp[key] = node

	if len(mp) > n {
		hd := getHead()
		remove(hd)
		delete(mp, hd.key)
	}
}

func get(key int) int {
	var node *DLL
	if _, ok := mp[key]; ok {
		node = mp[key]
	}
	if node == nil {
		return -1
	}
	remove(node)
	add(node)
	return node.data
}

func main() {
	head.next = tail
	tail.prev = head
	set(12, 30)
	set(13, 40)
	set(15, 50)
	set(16, 60)
	set(17, 70)
	set(18, 80)
	fmt.Println(get(12))
	fmt.Println(get(13))
	fmt.Println(get(14))
	fmt.Println(get(15))
	fmt.Println(get(16))
	fmt.Println(get(17))
	fmt.Println(get(18))
}
