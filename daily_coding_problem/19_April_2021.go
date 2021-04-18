// An XOR linked list is a more memory efficient doubly linked list.
// Instead of each node holding next and prev fields, it holds a field named both,
// which is an XOR of the next node and the previous node. Implement an XOR linked list;
// it has an add(element) which adds the element to the end, and a get(index) which returns the node at index.

// If using a language that has no pointers (such as Python),
// you can assume you have access to get_pointer and dereference_pointer
// functions that converts between nodes and memory addresses.

package main

import (
	"fmt"
	"unsafe"
)

type XORList struct {
	data interface{}
	both *XORList
}

func XOR(a, b *XORList) *XORList {
	x := uintptr(unsafe.Pointer(a)) ^ uintptr(unsafe.Pointer(b))
	y := (*XORList)(unsafe.Pointer(x))
	return y
}

func printList(head *XORList) {
	var prev *XORList = nil
	var curr *XORList = head
	var next *XORList
	fmt.Print("XORlist data : ")
	for curr != nil {
		fmt.Print(curr.data, " ")
		next = XOR(prev, curr.both)
		prev = curr
		curr = next
	}
	fmt.Println()
}

func add(head **XORList, data interface{}) {
	newNode := &XORList{data: data}
	if *head == nil {
		newNode.both = *head
		*head = newNode
	} else {
		var prev *XORList = nil
		var next *XORList
		var curr *XORList = *head
		for curr != nil {
			next = XOR(prev, curr.both)
			prev = curr
			curr = next
		}
		prev.both = XOR(prev.both, newNode)
		newNode.both = XOR(prev, nil)
	}
}

//get -> return the value at specific index
func get(head *XORList, index int) *XORList {
	var curr *XORList = head
	var prev *XORList = nil
	var next *XORList
	if head == nil {
		fmt.Println("xor list is empty")
		return nil
	} else {
		for curr != nil && index != 0 {
			next = XOR(prev, curr.both)
			prev = curr
			curr = next
			index--
		}
		if index == 0 {
			return prev
		} else {
			fmt.Println("index out of bounds")
			return nil
		}
	}
}

func main() {
	var head *XORList
	add(&head, 10)
	add(&head, 20)
	add(&head, 30)
	add(&head, 40)
	printList(head)
	var val *XORList
	val = get(head, 2)
	if val != nil {
		fmt.Printf("index: %d, value: %d\n", 2, val.data)
	}
}
