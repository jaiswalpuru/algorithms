/*
Author : Puru Jaiswal
Reference : CLRS
*/

package sll

import "fmt"

type Item interface{}

type Node struct {
	data Item
	next *Node
}

type List struct {
	head *Node
}

func New() *List { return &List{} }

func (l *List) Search(val interface{}) *Node {
	itr := l.head

	for itr != nil {
		if itr.data == val {
			return itr
		}
		itr = itr.next
	}

	return nil
}

func (l *List) Insert(node *Node) {
	if l == nil {
		l.head = node
	} else {
		temp := l.head
		node.next = temp
		l.head = node
	}
}

func (l *List) Delete(node *Node) {
	if l == nil {
		fmt.Println("List is empty")
	} else {
		itr := l.head
		//head is the node to be deleted
		if itr == node {
			l.head = l.head.next
		} else {
			var prev *Node = nil
			for itr != nil {
				if itr == node {
					prev.next = itr.next
					break
				}
				prev = itr
				itr = itr.next
			}
		}
	}
}
