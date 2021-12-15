/*
Implement locking in a binary tree.
A binary tree node can be locked or unlocked only if all of its descendants or ancestors are not locked.

Design a binary tree node class with the following methods:

•is_locked, which returns whether the node is locked

•lock, which attempts to lock the node. If it cannot be locked, then it should return false. Otherwise,
	it should lock it and return true.

• unlock, which unlocks the node. If it cannot be unlocked, then it should return false. Otherwise,
	it should unlock it and return true.

You may augment the node to add parent pointers or any other property you would like.
You may assume the class is used in a single-threaded program,
so there is no need for actual locks or mutexes. Each method should run in O(h), where h is the height of the tree.
*/

package main

import "fmt"

type Node struct {
	data                     int
	left                     *Node
	right                    *Node
	parent                   *Node
	is_locked                bool
	locked_descendants_count int
}

func NewNode(data int, parent *Node) *Node {
	return &Node{data: data, locked_descendants_count: 0, parent: parent}
}

func islock(node *Node) bool { return node.is_locked } //return true if the given node is locked

func lock(node *Node) bool {
	//It is possible only when ancetors and descendants of the current node are not locked
	if node.is_locked {
		return false
	}
	if !(can_lock_unlock(node)) {
		return false
	}

	node.is_locked = true
	cur := node.parent
	for cur != nil {
		cur.locked_descendants_count++
		cur = cur.parent
	}
	return true
}

func unlocks(node *Node) bool {
	//unlocks the node and update the information
	if !node.is_locked {
		return false //already locked
	}
	if can_lock_unlock(node) {
		return false
	}
	node.is_locked = false
	curr := node.parent
	for curr != nil {
		curr.locked_descendants_count--
		curr = curr.parent
	}
	return true
}

//checks whether the ancestors are locked or unlocked
func can_lock_unlock(node *Node) bool {
	if node.locked_descendants_count > 0 {
		return false
	}
	cur := node.parent
	for cur != nil {
		if cur.locked_descendants_count > 0 {
			return false
		}
		cur = cur.parent
	}
	return true
}

func main() {
	node := NewNode(12, nil)
	node.left = NewNode(10, node)
	node.right = NewNode(20, node)
	node.left.left = NewNode(5, node.left)
	node.left.right = NewNode(11, node.left)
	node.right.left = NewNode(15, node.right)
	node.right.right = NewNode(25, node.right)
	fmt.Println(lock(node.left.left))
	fmt.Println(lock(node.left))
	fmt.Println(unlocks(node.left.left))
}
