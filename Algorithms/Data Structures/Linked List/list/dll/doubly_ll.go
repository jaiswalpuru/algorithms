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

type List struct {
	head *Node
}

func New() *List { return &List{} }

func (l *List) Search(val Item) *Node {
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
	itr := l.head
	if itr == nil {
		l.head = node
	} else {
		node.next = itr
		itr.prev = node
		l.head = node
	}
}

//to delete a node from the list, first search should be called from the wrapper
func (l *List) Delete(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}
