/*
Author : Puru Jaiswal
Reference : CLRS
*/

package dll

type Item interface{}

type Node struct {
	data Item
	next *Node
	prev *Node
}

func New(data Item) *Node { return &Node{data: data, next: nil, prev: nil} }

func (l *Node) Search(val Item) *Node {
	for l.next != l {
		if l.data == val {
			return l
		}
		l = l.next
	}
	return nil
}

func (l *Node) Insert(n *Node) {
	n.next = l
	if l != nil {
		l.prev = n
	}
	l = n
	n.prev = nil
}

//to delete a node from the list, first search should be called from the wrapper
func (l *Node) Delete(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}
